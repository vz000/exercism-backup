package gross

func unitInUnits(units map[string]int, unit string) bool {
    var found bool = false
	for k := range units {
        if k == unit {
            found = true
        }
    }
    return found
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int {
    	"dozen": 12,
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    found := unitInUnits(units, unit)

    if found {
        if value, exists := bill[item]; exists {
            bill[item] = value + units[unit]
        } else {
            bill[item] = units[unit]
        }
    }
    return found
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    var result bool = false
	found := unitInUnits(units, unit)
    quantity, inBill := GetItem(bill, item)

    if inBill && found {
        resultQuantity := quantity - units[unit]
        if resultQuantity < 0 {
            result = false
        } else if resultQuantity == 0 {
            delete(bill, item)
            result = true
        } else {
            bill[item] = resultQuantity
            result = true
        }
    }
    
    return result
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    var inBill bool = false
    var quantity int = 0
    if _, exists := bill[item]; exists {
        inBill = true
        quantity = bill[item]
    }
    return quantity, inBill
}
