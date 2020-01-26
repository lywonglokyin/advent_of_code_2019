#include <iostream>

#include "../include/utils.h"

int main() {
    std::string command_text = readFile("input/day2")[0];
    std::vector<std::string> commands = split(command_text, ",");

    commands[1] = "12";
    commands[2] = "2";

    std::vector<int> commands_int;
    for (std::string command : commands) {
        commands_int.push_back(std::stoi(command));
    }
    unsigned int pointer = 0;
    bool exit_flag = false;
    while (!exit_flag) {
        int opcode = commands_int[pointer];
        int operand1 = commands_int[pointer + 1];
        int operand2 = commands_int[pointer + 2];
        int operand3 = commands_int[pointer + 3];
        switch (opcode) {
            case 1:
                commands_int[operand3] =
                    commands_int[operand1] + commands_int[operand2];
                break;
            case 2:
                commands_int[operand3] =
                    commands_int[operand1] * commands_int[operand2];
                break;
            case 99:
                exit_flag = true;
                break;
        }
        pointer += 4;
    }
    std::cout << commands_int[0] << std::endl;
}