package sum

func sum(a, b int) int {
	return a + b
}

func sums(xs ...int) int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}
