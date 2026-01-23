package purchase
import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
        return true
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    finalChoice := option2
    if option1 < option2 {
        finalChoice = option1
    }
    return fmt.Sprintf("%s is clearly the better choice.", finalChoice)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    finalCost := originalPrice
	if age < 3 {
        finalCost *= 0.80
    } else if age >= 10 {
        finalCost *= 0.50
    } else if age >= 3 && age < 10 {
        finalCost *= 0.70
    }
    return finalCost
}
