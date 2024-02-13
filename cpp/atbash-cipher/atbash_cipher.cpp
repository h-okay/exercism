#include "atbash_cipher.h"
#include <string>
#include <iostream>

namespace atbash_cipher
{
    char invert(char c)
    {
        return std::isdigit(c) ? c : (char)('z' - (std::tolower(c) - 'a'));
    }

    std::string encode(std::string message)
    {
        std::string encoded;
        int char_count = 0;
        for (const char &c : message)
        {
            if (isalnum(c))
            {
                if (char_count == 5)
                {
                    encoded += ' ';
                    char_count = 0;
                }
                encoded += invert(c);
                ++char_count;
            }
        }
        return encoded;
    }

    std::string decode(std::string message)
    {
        std::string decoded;
        for (const char &c : message)
        {
            if (std::isalnum(c)) decoded += invert(c);
        }
        return decoded;
    }

} // namespace atbash_cipher
