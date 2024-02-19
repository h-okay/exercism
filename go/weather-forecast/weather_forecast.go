// Package weather provides functions
// to format weather outputs.
package weather

// CurrentCondition describes current weather condition.
var CurrentCondition string

// CurrentLocation describes current location.
var CurrentLocation string

// Forecast returns a formatted string with the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
