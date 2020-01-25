#include <fstream>
#include <iostream>
#include <string>
#include <vector>

#include "../include/utils.h"

std::vector<std::string> split(const std::string &s,
                               const std::string &delimiter) {
    std::vector<std::string> v;

    size_t last = 0;
    size_t next = 0;
    while ((next = s.find(delimiter, last)) != std::string::npos) {
        v.push_back(s.substr(last, next - last));
        last = next + 1;
    }
    v.push_back(s.substr(last));
    return v;
};

std::vector<std::string> readFile(const std::string &filepath) {
    std::vector<std::string> lines;
    std::string line;

    std::ifstream file(filepath);

    if (file.is_open()) {
        while (getline(file, line)) {
            lines.push_back(line);
        }
        file.close();
    }
    return lines;
}