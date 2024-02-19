#include "hexadecimal.h"
#include <cmath>
#include <map>

namespace hexadecimal
{
    int convert(std::string input)
    {
        std::map<char, int> mapping{{'a', 10}, {'b', 11}, {'c', 12}, {'d', 13}, {'e', 14}, {'f', 15}};
        int decimal{0};
        int multiplier{0};
        for (int i = static_cast<int>(input.length() - 1); i >= 0; i--)
        {
            char ch = input[i];
            if (isalpha(ch))
            {
                auto value = mapping.find(tolower(ch))->second;
                if (!value)
                    return 0;
                decimal += value * std::pow(16, multiplier);
            }
            else if (isdigit(ch))
            {
                decimal += (ch - '0') * std::pow(16, multiplier);
            }
            multiplier++;
        }
        return decimal;
    }
} // namespace hexadecimal
