package purchase

const (
    VehicleTypeCar = "car"
    VehicleTypeTruck = "truck"

    ResellPercentUnder3Years  = 80.0
	ResellPercentUnder10Years = 70.0
	ResellPercentOver10Years  = 50.0
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == VehicleTypeCar || kind == VehicleTypeTruck
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		return option2 + " is clearly the better choice."
	}
	return option1 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	retainedPercent := ResellPercentOver10Years
    if age < 3 {
        retainedPercent = ResellPercentUnder3Years
    } else if age < 10 {
        retainedPercent = ResellPercentUnder10Years
    }
    return originalPrice / 100 * retainedPercent
}
