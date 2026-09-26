package supermarketcoupon

func logic(n int, p []int) int {
	if n == 0 {
		return 0
	}

	topIndex := 0
	for i := range p {
		if p[i] > p[topIndex] {
			topIndex = i
		}
	}

	p[topIndex] = p[topIndex] / 2

	var sum int
	for _, v := range p {
		sum += v
	}

	return sum
}
