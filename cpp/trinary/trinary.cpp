#include "trinary.h"
#include <iostream>
#include <cmath>

namespace trinary
{
    int to_decimal(std::string input)
    {
        int decimal{0};
        int multiplier{0};
        for (int i = static_cast<int>(input.length() - 1); i >= 0; i--)
        {
            if (isdigit(input[i]))
            {
                decimal += (input[i] - '0') * std::pow(3, multiplier);
                multiplier++;
            }
        }
        return decimal;
    }
} // namespace trinary
