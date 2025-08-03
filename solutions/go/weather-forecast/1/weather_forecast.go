// Package weather provides tools to forecast the weather of various cities in Goblinocus.
package weather

// CurrentCondition represents the current weather conditions in a city.
var CurrentCondition string
// CurrentLocation represents the city whose weather we want to know.
var CurrentLocation string

// Forecast returns a string with a message that tells the current weather conditions of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
