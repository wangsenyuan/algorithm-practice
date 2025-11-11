package p2169

func countOperations(num1 int, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}
	// a - b
	var res int

	for num2 > 0 {
		res += num1 / num2
		num1, num2 = num2, num1%num2
	}

	return res
}
