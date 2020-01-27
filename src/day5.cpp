#include <iostream>

#include "../include/utils.h"

std::string pad_zero(const std::string& s, int length = 5);

int main() {
    std::string command_text = readFile("input/day5")[0];
    std::vector<std::string> commands = split(command_text, ",");

    unsigned int pointer = 0;
    bool exit_flag = false;
    while (!exit_flag) {
        std::string param = pad_zero(commands[pointer]);
        int opcode = std::stoi(param.substr(3));
        switch (opcode) {
            case 1: {
                if (param[0] == '1') {
                    pointer += 4;
                    break;
                }
                int operand1 = std::stoi(commands[pointer + 1]);
                int operand2 = std::stoi(commands[pointer + 2]);
                int operand3 = std::stoi(commands[pointer + 3]);
                operand1 =
                    param[2] == '1' ? operand1 : std::stoi(commands[operand1]);
                operand2 =
                    param[1] == '1' ? operand2 : std::stoi(commands[operand2]);
                commands[operand3] = std::to_string(operand1 + operand2);
                pointer += 4;
                break;
            }
            case 2: {
                if (param[0] == '1') {
                    pointer += 4;
                    break;
                }
                int operand1 = std::stoi(commands[pointer + 1]);
                int operand2 = std::stoi(commands[pointer + 2]);
                int operand3 = std::stoi(commands[pointer + 3]);
                operand1 =
                    param[2] == '1' ? operand1 : std::stoi(commands[operand1]);
                operand2 =
                    param[1] == '1' ? operand2 : std::stoi(commands[operand2]);
                commands[operand3] = std::to_string(operand1 * operand2);
                pointer += 4;
                break;
            }
            case 3: {
                int operand1 = std::stoi(commands[pointer + 1]);
                std::string value;
                std::cin >> value;
                commands[operand1] = value;
                pointer += 2;
                break;
            }
            case 4: {
                int operand1 = std::stoi(commands[pointer + 1]);
                std::cout << commands[operand1] << std::endl;
                pointer += 2;
                break;
            }
            case 5: {
                int operand1 = std::stoi(commands[pointer + 1]);
                operand1 =
                    param[2] == '1' ? operand1 : std::stoi(commands[operand1]);

                if (operand1 != 0) {
                    int operand2 = std::stoi(commands[pointer + 2]);
                    operand2 = param[1] == '1' ? operand2
                                               : std::stoi(commands[operand2]);
                    pointer = operand2;
                } else {
                    pointer += 3;
                }
                break;
            }
            case 6: {
                int operand1 = std::stoi(commands[pointer + 1]);
                operand1 =
                    param[2] == '1' ? operand1 : std::stoi(commands[operand1]);

                if (operand1 == 0) {
                    int operand2 = std::stoi(commands[pointer + 2]);
                    operand2 = param[1] == '1' ? operand2
                                               : std::stoi(commands[operand2]);
                    pointer = operand2;
                } else {
                    pointer += 3;
                }
                break;
            }
            case 7: {
                int operand1 = std::stoi(commands[pointer + 1]);
                int operand2 = std::stoi(commands[pointer + 2]);
                int operand3 = std::stoi(commands[pointer + 3]);
                operand1 =
                    param[2] == '1' ? operand1 : std::stoi(commands[operand1]);
                operand2 =
                    param[1] == '1' ? operand2 : std::stoi(commands[operand2]);
                if (param[0] == '1') {
                    pointer += 4;
                    break;
                }
                if (operand1 < operand2) {
                    commands[operand3] = "1";
                } else {
                    commands[operand3] = "0";
                }
                pointer += 4;
                break;
            }
            case 8: {
                int operand1 = std::stoi(commands[pointer + 1]);
                int operand2 = std::stoi(commands[pointer + 2]);
                int operand3 = std::stoi(commands[pointer + 3]);
                operand1 =
                    param[2] == '1' ? operand1 : std::stoi(commands[operand1]);
                operand2 =
                    param[1] == '1' ? operand2 : std::stoi(commands[operand2]);
                if (param[0] == '1') {
                    pointer += 4;
                    break;
                }
                if (operand1 == operand2) {
                    commands[operand3] = "1";
                } else {
                    commands[operand3] = "0";
                }
                pointer += 4;
                break;
            }
            case 99: {
                exit_flag = true;
                break;
            }
        }
    }
}

std::string pad_zero(const std::string& s, int length) {
    std::string temp = s;
    for (int i = 0; i < length - s.length(); ++i) {
        temp = "0" + temp;
    }
    return temp;
};