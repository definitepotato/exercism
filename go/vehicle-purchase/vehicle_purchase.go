package purchase

import (
	"fmt"
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	vehicles := []string{}
	vehicles = append(vehicles, option1)
	vehicles = append(vehicles, option2)

	sort.Strings(vehicles)

	return fmt.Sprintf("%s is clearly the better choice.", vehicles[0])
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return (80.00 / 100.00) * originalPrice
	}

	if age >= 3 && age < 10 {
		return (70.00 / 100.00) * originalPrice
	}

	if age >= 10 {
		return (50.00 / 100.00) * originalPrice
	}

	return 0.0
}
