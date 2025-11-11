package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
	if tu == Celsius {
		return "°C"
	}

	return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%s %s", fmt.Sprint(t.degree), t.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string {
	if su == KmPerHour {
		return "km/h"
	}

	return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%s %s", fmt.Sprint(s.magnitude), s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (md *MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s %s, Wind %s at %s %s, %s%% Humidity", md.location, fmt.Sprint(md.temperature.degree), md.temperature.unit.String(), md.windDirection, fmt.Sprint(md.windSpeed.magnitude), md.windSpeed.unit.String(), fmt.Sprint(md.humidity))
}
