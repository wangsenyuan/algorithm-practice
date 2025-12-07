package p3765

func completePrime(num int) bool {
	var right int
	base := 1

	checkPrime := func(num int) bool {
		if num == 1 {
			return false
		}
		for i := 2; i <= num/i; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	if !checkPrime(num) {
		return false
	}

	for num > 0 {
		x := num % 10
		right += x * base
		base *= 10
		num /= 10
		if !checkPrime(right) || num > 0 && !checkPrime(num) {
			return false
		}
	}

	return true
}
