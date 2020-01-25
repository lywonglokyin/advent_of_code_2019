#include <iostream>
#include <vector>

#include "../include/utils.h"

int max(int a, int b) { return a > b ? a : b; }
int min(int a, int b) { return a > b ? b : a; }

struct Point {
    int x;
    int y;
};

struct Line {
    Point p1;
    Point p2;

    int length() {
        return max(p1.x, p2.x) - min(p1.x, p2.x) + max(p1.y, p2.y) -
               min(p1.y, p2.y);
    }

    bool contains_point(const Point& p) {
        if (p1.x == p2.x) {
            return (p.x == p1.x) && (p.y <= max(p1.y, p2.y)) &&
                   (p.y >= min(p1.y, p2.y));
        } else {
            return (p.y == p1.y) && (p.x <= max(p1.x, p2.x)) &&
                   (p.y >= min(p1.x, p2.x));
        }
    }
};

// If multiple intercept excist, then return the one closest to (0, 0)
// For now, it assumes the lines are either horizontal or vertical
Point intercept(const Line& l1, const Line& l2);

bool hasIntercept(const Line& l1, const Line& l2);

bool isHorizontal(const Line& line);

std::vector<Line> parse_command(const std::vector<std::string>& commands);

int manhattan_distance(const Point& p);

int calculate_delay(const Point& p, const std::vector<Line>& segments);

int main() {
    Point icpt = intercept(Line{{1, -1}, {1, 5}}, Line{{1, -5}, {1, 6}});
    std::vector<std::string> inputs = readFile("input/day3");
    std::vector<std::string> pipe_1_commands = split(inputs[0], ",");
    std::vector<std::string> pipe_2_commands = split(inputs[1], ",");
    std::vector<Line> pipe_1_lines = parse_command(pipe_1_commands);
    std::vector<Line> pipe_2_lines = parse_command(pipe_2_commands);

    std::vector<Point> intercepts;
    for (Line p1_line : pipe_1_lines) {
        for (Line p2_line : pipe_2_lines) {
            if (hasIntercept(p1_line, p2_line)) {
                intercepts.push_back(intercept(p1_line, p2_line));
            }
        }
    }

    int shortest = INT32_MAX;
    for (Point incpt : intercepts) {
        int distance = manhattan_distance(incpt);
        if (distance == 0) {  // cheap way to ignore first intersection at (0,
                              // 0)
            continue;
        }
        if (distance < shortest) {
            shortest = distance;
        }
    }
    std::cout << shortest << std::endl;

    int shortest_delay = INT32_MAX;
    for (Point incpt : intercepts) {
        if (incpt.x == 0 && incpt.y == 0) {
            continue;
        }
        int p1_delay = calculate_delay(incpt, pipe_1_lines);
        int p2_delay = calculate_delay(incpt, pipe_2_lines);
        int total_delay = p1_delay + p2_delay;
        if (total_delay < shortest_delay) {
            shortest_delay = total_delay;
        }
    }
    std::cout << shortest_delay << std::endl;
    return 0;
}

bool hasIntercept(const Line& l1, const Line& l2) {
    int l1_left_x = min(l1.p1.x, l1.p2.x);
    int l1_right_x = max(l1.p1.x, l1.p2.x);
    int l1_lower_y = min(l1.p1.y, l1.p2.y);
    int l1_upper_y = max(l1.p1.y, l1.p2.y);
    int l2_left_x = min(l2.p1.x, l2.p2.x);
    int l2_right_x = max(l2.p1.x, l2.p2.x);
    int l2_lower_y = min(l2.p1.y, l2.p2.y);
    int l2_upper_y = max(l2.p1.y, l2.p2.y);
    if ((l1_left_x <= l2_right_x) && (l1_right_x >= l2_left_x)) {
        if ((l1_lower_y <= l2_upper_y) && (l1_upper_y >= l2_lower_y)) {
            return true;
        }
    }
    return false;
}

Point intercept(const Line& l1, const Line& l2) {
    bool l1_horizontal = isHorizontal(l1);
    bool l2_horizontal = isHorizontal(l2);
    int l1_left_x = min(l1.p1.x, l1.p2.x);
    int l1_right_x = max(l1.p1.x, l1.p2.x);
    int l1_lower_y = min(l1.p1.y, l1.p2.y);
    int l1_upper_y = max(l1.p1.y, l1.p2.y);
    int l2_left_x = min(l2.p1.x, l2.p2.x);
    int l2_right_x = max(l2.p1.x, l2.p2.x);
    int l2_lower_y = min(l2.p1.y, l2.p2.y);
    int l2_upper_y = max(l2.p1.y, l2.p2.y);
    if (l1_horizontal && l2_horizontal) {
        int smallX = max(l1_left_x, l2_left_x);
        int bigX = min(l1_right_x, l2_right_x);
        if (smallX > 0) {
            return {smallX, l1_lower_y};
        } else if (bigX < 0) {
            return {bigX, l1_lower_y};
        } else {
            return {0, l1_lower_y};
        }
    } else if ((!l1_horizontal) && (!l2_horizontal)) {
        int smallY = max(l1_lower_y, l2_lower_y);
        int bigY = min(l1_upper_y, l2_upper_y);
        if (smallY > 0) {
            return {l1_left_x, smallY};
        } else if (bigY < 0) {
            return {l1_left_x, bigY};
        } else {
            return {l1_left_x, 0};
        }
    } else {
        if (l1_horizontal) {
            return {l2_left_x, l1_lower_y};
        } else {
            return {l1_left_x, l2_lower_y};
        }
    }
}

bool isHorizontal(const Line& line) { return (line.p1.y == line.p2.y); }

std::vector<Line> parse_command(const std::vector<std::string>& commands) {
    std::vector<Line> segments;
    Point position{0, 0};
    for (std::string command : commands) {
        char direction = command[0];
        int displacement = stoi(command.substr(1));
        Point new_position{position};
        switch (direction) {
            case 'U':
                new_position.y += displacement;
                break;
            case 'D':
                new_position.y -= displacement;
                break;
            case 'L':
                new_position.x -= displacement;
                break;
            case 'R':
                new_position.x += displacement;
                break;
        }
        segments.push_back(Line{position, new_position});
        position = new_position;
    }
    return segments;
}

int manhattan_distance(const Point& p) {
    int x = p.x >= 0 ? p.x : -p.x;
    int y = p.y >= 0 ? p.y : -p.y;
    return x + y;
}

int calculate_delay(const Point& p, const std::vector<Line>& segments) {
    int delay = 0;
    for (Line segment : segments) {
        if (segment.contains_point(p)) {
            if (isHorizontal(segment)) {
                int displace = (p.x - segment.p1.x);
                delay += (displace > 0 ? displace : -displace);
                return delay;
            } else {
                int displace = (p.y - segment.p1.y);
                delay += (displace > 0 ? displace : -displace);
                return delay;
            }
        } else {
            delay += segment.length();
        }
    }
}