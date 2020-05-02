package main

// Rate returns string based on rating number.
func Rate(rate int) string {
	switch rate {
	case 0, 1:
		return "Worst"
	case 2:
		return "Ok"
	case 3:
		return "Average"
	case 4:
		return "Above Average"
	case 5:
		return "Superb"
	default:
		return "Invalid rating"
	}
}
