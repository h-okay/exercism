#include "reverse_string.h"

namespace reverse_string
{
    std::string reverse_string(std::string str)
    {
        int left{0};
        int right{(int) str.length() - 1};
        char temp;
        while (left < right)
        {
            temp = str[left];
            str[left] = str[right];
            str[right] = temp;
            left++;
            right--;
        }
        return str;
    }
} // namespace reverse_string
