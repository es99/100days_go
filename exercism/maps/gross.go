package main

func main() {}

func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

func newBill() map[string]int {
	newBill := make(map[string]int)
	return newBill
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	for k, v := range units {
		if k == unit {
			bill[item] += v
			return true
		}
	}
	return false
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	_, existsBill := bill[item]
	if exists && existsBill {
		newQtd := bill[item] - units[unit]
		if newQtd < 0 {
			return false
		} else if newQtd == 0 {
			delete(bill, item)
			return true
		} else {
			bill[item] -= units[unit]
			return true
		}
	}
	return false
}

func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	if exists {
		return bill[item], true
	} else {
		return value, false
	}
}
