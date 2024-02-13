#include "rotational_cipher.h"
#include <iostream>

namespace rotational_cipher
{
    std::string rotate(std::string input, int rotation)
    {
        std::string characters = "abcdefghijklmnopqrstuvwxyz";
        std::string encoded;
        for (const char &ch : input)
        {
            if (std::isalpha(ch))
            {
                if (std::islower(ch))
                {
                    encoded += characters[(characters.find(ch) + rotation) % 26];
                    continue;
                }
                if (std::isupper(ch))
                {
                    encoded += std::toupper(characters[(characters.find(std::tolower(ch)) + rotation) % 26]);
                    continue;
                }
            }
            encoded += ch; // not alphanum
        }
        return encoded;
    }
} // namespace rotational_cipher
