package ticket

func Price(age int) float64 {
	if age <= 3 {
		return 0
	} else if age <= 15 {
		return 15
	} else if age <= 50 {
		return 30
	}
	return 5

}
