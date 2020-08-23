#include <iostream>
#include <string>

bool hasAdjacentDigit(int number);
bool increasingDigits(int number);
bool exactPairDigits(int number);

int main() {
    const int START_NUM = 367479;
    const int END_NUM = 893698;

    int count = 0;

    for (int i = START_NUM; i <= END_NUM; ++i) {
        if (!hasAdjacentDigit(i)) {
            continue;
        }
        if (!increasingDigits(i)) {
            continue;
        }
        ++count;
    }
    std::cout << count << std::endl;

    int part_two_count = 0;
    for (int i = START_NUM; i <= END_NUM; ++i) {
        if (!hasAdjacentDigit(i)) {
            continue;
        }
        if (!increasingDigits(i)) {
            continue;
        }
        if (!exactPairDigits(i)) {
            continue;
        }
        ++part_two_count;
    }
    std::cout << part_two_count << std::endl;
}

bool hasAdjacentDigit(int number) {
    std::string number_str = std::to_string(number);
    for (unsigned int i = 0; i < number_str.length() - 1; ++i) {
        if (number_str[i] == number_str[i + 1]) {
            return true;
        }
    }
    return false;
}

bool increasingDigits(int number) {
    std::string number_str = std::to_string(number);
    int digit = number_str[0] - '0';
    for (unsigned int i = 1; i < number_str.length(); ++i) {
        if ((number_str[i] - '0') < digit) {
            return false;
        }
        digit = number_str[i] - '0';
    }
    return true;
}

bool exactPairDigits(int number) {
    std::string number_str = std::to_string(number);
    int length = 1;
    for (unsigned int i = 0; i < number_str.length() - 1; ++i) {
        if (number_str[i] == number_str[i + 1]) {
            ++length;
        } else {
            if (length == 2) {
                return true;
            } else {
                length = 1;
            }
        }
    }
    return (length == 2);
}