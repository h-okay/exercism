#include "darts.h"

namespace darts
{
    int score(float x, float y)
    {
        if (x * x + y * y <= 1)
        {
            return 10;
        }
        if (x * x + y * y <= 5 * 5)
        {
            return 5;
        }
        if (x * x + y * y <= 10 * 10)
        {
            return 1;
        }
        return 0;
    }
} // namespace darts
