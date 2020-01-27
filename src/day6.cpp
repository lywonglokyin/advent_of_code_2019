#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>

#include "../include/utils.h"

struct OrbitNode {
    std::string name;
    std::vector<OrbitNode*> childOribitNode;

    OrbitNode() = delete;
    OrbitNode(std::string name) : name{name} {}
    OrbitNode(const OrbitNode& another) = delete;

    void addChild(OrbitNode* child) { childOribitNode.push_back(child); }

    ~OrbitNode() {
        for (OrbitNode* child : childOribitNode) {
            delete child;
        }
    }
};

int countDirectAndIndirectOrbit(const OrbitNode* node, int level = 0);

std::vector<std::string> getChainOfOrbits(const OrbitNode* node,
                                          const std::string& name);

// Assumption:
// 1) No orbits have more than 1 parent orbits.
// 2) Same parent-child pair won't reappear.
int main() {
    std::unordered_map<std::string, OrbitNode*> nodes;
    std::unordered_map<std::string, OrbitNode*> roots;

    std::vector<std::string> orbit_infos = readFile("input/day6");

    for (auto orbit_info : orbit_infos) {
        std::vector<std::string> orbits = split(orbit_info, ")");
        std::string parent_orbit = orbits[0];
        std::string child_orbit = orbits[1];

        auto find_parent = nodes.find(parent_orbit);
        auto find_child = nodes.find(child_orbit);
        if ((find_parent == nodes.end()) && (find_child == nodes.end())) {
            // Parent and child not exist yet.
            OrbitNode* parent = new OrbitNode(parent_orbit);
            OrbitNode* child = new OrbitNode(child_orbit);
            parent->addChild(child);
            nodes[parent_orbit] = parent;
            nodes[child_orbit] = child;
            roots[parent_orbit] = parent;
        } else if (find_parent == nodes.end()) {
            // Parent not exist, child exist
            OrbitNode* parent = new OrbitNode(parent_orbit);
            parent->addChild(find_child->second);
            // Below is based on assumption 1
            nodes[parent_orbit] = parent;
            roots.erase(child_orbit);
            roots[parent_orbit] = parent;
        } else if (find_child == nodes.end()) {
            // Parent exist, child not exist.
            OrbitNode* child = new OrbitNode(child_orbit);
            find_parent->second->addChild(child);
            nodes[child_orbit] = child;
        } else {
            // Both exist
            // Also assumption 1
            // Parent not exist, child exist
            find_parent->second->addChild(find_child->second);
            // Below is based on assumption 1
            roots.erase(child_orbit);
        }
    }

    std::cout << "No. of roots: " << roots.size() << std::endl;
    int count = 0;
    for (auto it : roots) {
        count += countDirectAndIndirectOrbit(it.second);
    }
    std::cout << "Count: " << count << std::endl;

    for (auto it : roots) {
        auto chain_YOU = getChainOfOrbits(it.second, "YOU");
        auto chain_SAN = getChainOfOrbits(it.second, "SAN");
        const int PTR_YOU_END = chain_YOU.size() - 1;
        const int PTR_SAN_END = chain_SAN.size() - 1;
        int i = 0;
        while (true) {
            if (chain_SAN[PTR_SAN_END - i] != chain_YOU[PTR_YOU_END - i]) {
                break;
            }
            ++i;
        }
        std::cout << "Need to jump: "
                  << chain_YOU.size() - i - 1 + chain_SAN.size() - i - 1
                  << std::endl;
    }

    for (auto it : roots) {
        delete it.second;
    }
}

int countDirectAndIndirectOrbit(const OrbitNode* node, int level) {
    int count = 0;
    for (OrbitNode* child : node->childOribitNode) {
        ++count;
        count += level;
        count += countDirectAndIndirectOrbit(child, level + 1);
    }
    return count;
}

bool dfs_recursive(const OrbitNode* node, const std::string& name,
                   std::vector<std::string>& chain) {
    if (node->name == name) {
        chain.push_back(node->name);
        return true;
    }
    for (OrbitNode* child : node->childOribitNode) {
        if (dfs_recursive(child, name, chain)) {
            chain.push_back(node->name);
            return true;
        }
    }
    return false;
}

std::vector<std::string> getChainOfOrbits(const OrbitNode* node,
                                          const std::string& name) {
    std::vector<std::string> chain;
    dfs_recursive(node, name, chain);
    return chain;
}