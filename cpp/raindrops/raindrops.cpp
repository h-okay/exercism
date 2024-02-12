#include "raindrops.h"
#include <string>
#include <iostream>

using namespace std;

namespace raindrops
{
    bool divisible_by_3(int number)
    {
        return number % 3 == 0;
    }
    bool divisible_by_5(int number)
    {
        return number % 5 == 0;
    }
    bool divisible_by_7(int number)
    {
        return number % 7 == 0;
    }
    string convert(int number)
    {
        string output = "";
        if (divisible_by_3(number))
        {
            output += "Pling";
        }
        if (divisible_by_5(number))
        {
            output += "Plang";
        }
        if (divisible_by_7(number))
        {
            output += "Plong";
        }
        if (output.length() == 0)
        {
            return to_string(number);
        }
        return output;
    }

} // namespace raindrops
