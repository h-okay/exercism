// INFO: Headers from the standard library should be inserted at the top via
// #include <LIBRARY_NAME>
#include <cmath>

const int DAYS_IN_MONTH = 22;

// daily_rate calculates the daily rate given an hourly rate
double daily_rate(double hourly_rate)
{
    return hourly_rate * 8;
}

// apply_discount calculates the price after a discount
double apply_discount(double before_discount, double discount)
{
    return before_discount - (before_discount * discount * 0.01);
}

// monthly_rate calculates the monthly rate, given an hourly rate and a discount
// The returned monthly rate is rounded up to the nearest integer.
int monthly_rate(double hourly_rate, double discount)
{
    return ceil(apply_discount(daily_rate(hourly_rate) * DAYS_IN_MONTH, discount));
}

// days_in_budget calculates the number of workdays given a budget, hourly rate,
// and discount The returned number of days is rounded down (take the floor) to
// the next integer.
int days_in_budget(int budget, double hourly_rate, double discount)
{

    return std::floor(budget / apply_discount(daily_rate(hourly_rate), discount));
}
