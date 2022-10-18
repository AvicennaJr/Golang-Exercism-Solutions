 
// Package weather is a package that can forecast the current weather condition of various cities in Goblinocus.
package weather 

// CurrentCondition is a variable that stores current condition.
var CurrentCondition string 

// CurrentLocation is a variable that stores current location.
var CurrentLocation string 


// Forecast is a function returns the weather condition of a location.
func Forecast(city, condition string) string {
	 
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
