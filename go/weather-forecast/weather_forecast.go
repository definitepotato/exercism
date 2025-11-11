// Package weather provides weather information.
package weather

var (
	// CurrentCondition is  current weather conditions.
	CurrentCondition string
	// CurrentLocation is the name of a geographical location.
	CurrentLocation string
)

// Forecasts weather based on location and conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
