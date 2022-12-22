// Package weather provides tools to give a forecast of the weather.
package weather

// CurrentCondition represents the weather condition currently happening.
var CurrentCondition string
// CurrentLocation represents the location where the forecast takes place.
var CurrentLocation string

// Forecast returns a string with the current forecast using two string parameters.
// The first: city, representing the city location where the forecast takes place.
// The second: condition, representing the current effect in the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
