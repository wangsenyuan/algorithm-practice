package p3536

func maxProduct(n int) int {
	var first, second int
	for n > 0 {
		r := n % 10
		if r > first {
			second = first
			first = r
		} else if r > second {
			second = r
		}
		n /= 10
	}
	return first * second
}
