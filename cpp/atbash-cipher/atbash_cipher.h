#if !defined(ATBASH_CIPHER_H)
#define ATBASH_CIPHER_H
#include <string>

namespace atbash_cipher
{
    char invert(char c);
    std::string encode(std::string message);
    std::string decode(std::string message);
} // namespace atbash_cipher

#endif // ATBASH_CIPHER_H
