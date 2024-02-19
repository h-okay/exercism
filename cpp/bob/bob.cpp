#include "bob.h"
#include <algorithm>

namespace bob
{
    std::string hey(std::string input)
    {
        if (std::all_of(input.begin(), input.end(), [](char c)
                        { if (isalpha(c)) return std::isupper(c); return -1; }) and
            input[input.length() - 1] == '?')
        {
            return "Calm down, I know what I'm doing!";
        }
        if (std::all_of(input.begin(), input.end(), [](char c)
                        { if (isalpha(c)) return std::isupper(c); return -1; }))
        {
            return "Whoa, chill out!";
        }
        if (std::all_of(input.begin(), input.end(), [](char c)
                        { return std::isspace(c); }))
        {
            return "Fine. Be that way!";
        }
        if (input[input.length() - 1] == '?')
        {
            return "Sure.";
        }
        return "Whatever.";
    }
} // namespace bob
