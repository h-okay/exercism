#include <string>
#include <iostream>
using namespace std;

namespace log_line
{
    string message(string line)
    {
        return line.substr(line.find(":") + 2);
    }

    string log_level(string line)
    {
        return line.substr(1, line.find(":") - 2);
    }

    string reformat(string line)
    {
        string message = log_line::message(line);
        string level = log_line::log_level(line);
        return message + " (" + level + ")";
    }
} // namespace log_line
