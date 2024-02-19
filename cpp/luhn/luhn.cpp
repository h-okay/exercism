#include "luhn.h"
#include <algorithm>
#include <cctype>
#include <iostream>

namespace luhn
{
    bool valid(std::string number)
    {

        number.erase(std::remove_if(number.begin(), number.end(), ::isspace), number.end());
        if (number.length() <= 1) return false;

        int total{0};
        int digit{0};
        for (int i = number.length() - 1; i >= 0; i--)
        {
            if (!isdigit(number[i])) return false;

            int value = number[i] - '0';
            if ((digit + 1) % 2 == 0)
            {
                (value * 2 > 9) ? total += value * 2 - 9 : total += value * 2;
                digit++;
                continue;
            }

            total += value;
            digit++;
        }

        return total % 10 == 0;
    }
} // namespace luhn
