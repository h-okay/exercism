#include "vehicle_purchase.h"
#include <string>
#include <algorithm>
#include <vector>

using namespace std;
vector<string> LICENCE_REQUIRED{"car", "truck"};

namespace vehicle_purchase
{

    // needs_license determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
    bool needs_license(std::string kind)
    {
        vector<string>::iterator it = find(begin(LICENCE_REQUIRED), end(LICENCE_REQUIRED), kind);
        if (it != LICENCE_REQUIRED.end())
        {
            return true;
        }
        return false;
    }

    // choose_vehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
    std::string choose_vehicle(std::string option1, std::string option2)
    {
        string output = " is clearly the better choice.";
        if (option1 > option2)
        {
            return option2 + output;
        }
        return option1 + output;
    }

    // calculate_resell_price calculates how much a vehicle can resell for at a certain age.
    double calculate_resell_price(double original_price, double age)
    {
        if (age < 3)
        {
            return original_price * 0.8;
        }
        else if (age > 3 && age < 10)
        {
            return original_price * 0.7;
        }
        else if (age >= 10)
        {
            return original_price * 0.5;
        }
        else
        {
            return original_price;
        }
    }

} // namespace vehicle_purchase

/* alership. Since you are interested in buying a used vehicle, the price depends on how old the vehicle is. For a rough estimate, assume if the vehicle is
less than 3 years old, it costs 80% of the original price it had when it was brand new. If it is
at least 10 years old, it costs 50%. If the vehicle is
at least 3 years old but less than 10 years, it costs 70% of the original price. */
