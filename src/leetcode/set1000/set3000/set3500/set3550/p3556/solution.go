package p3556

func sumOfLargestPrimes(s string) int64 {
	n := len(s)
	var res [3]int64

	vis := make(map[int64]bool)

	vis[1] = true
	vis[0] = true

	for i := 0; i < n; i++ {
		var num int64
		for j := i; j < n; j++ {
			num = num*10 + int64(s[j]-'0')
			if vis[num] {
				continue
			}
			vis[num] = true
			if checkPrime(num) {
				if num > res[0] {
					res[2] = res[1]
					res[1] = res[0]
					res[0] = num
				} else if num > res[1] {
					res[2] = res[1]
					res[1] = num
				} else if num > res[2] {
					res[2] = num
				}
			}
		}
	}

	return res[0] + res[1] + res[2]
}

func checkPrime(num int64) bool {
	for x := int64(2); x <= num/x; x++ {
		if num%x == 0 {
			return false
		}
	}
	return true
}
