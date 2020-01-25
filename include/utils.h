#ifndef UTILS_H
#define UTILS_H

#include <string>
#include <vector>

std::vector<std::string> split(const std::string &s,
                               const std::string &delimiter);

std::vector<std::string> readFile(const std::string &filepath);

#endif