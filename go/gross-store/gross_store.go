package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}

	if val, ok := units[unit]; ok {
		bill[item] += val
		return true
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := bill[item]; !ok {
		return false
	}

	if _, ok := units[unit]; !ok {
		return false
	}

	if val, ok := units[unit]; ok {
		if bill[item]-val < 0 {
			return false
		}

		bill[item] -= val

		if bill[item] == 0 {
			delete(bill, item)
			return true
		}

		return true
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if val, ok := bill[item]; ok {
		return val, true
	}

	return 0, false
}
