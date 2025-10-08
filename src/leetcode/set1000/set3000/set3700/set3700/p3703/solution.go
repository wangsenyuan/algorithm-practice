package p3703

import "slices"

func countNoZeroPairs(n int64) int64 {
	// a + b = n
	// a, b中不包含0
	// dp[i][c] 表示后面i个数字满足条件，且进位为c(0,1)的个数
	var digits []int
	for num := n; num > 0; num /= 10 {
		digits = append(digits, int(num%10))
	}
	slices.Reverse(digits)

	m := len(digits)

	if m == 1 {
		// 3 = (1 + 2), (2 + 1)
		return max(0, int64(digits[0]-1))
	}

	// 没说不能leading zero

	dp := make([][][]int, m)
	for i := range m {
		dp[i] = make([][]int, 2)
		for j := range 2 {
			dp[i][j] = make([]int, 4)
			for x := range 4 {
				dp[i][j][x] = -1
			}
		}
	}

	var f func(i int, loan int, leading_zero int) int
	f = func(i int, loan int, leading_zero int) (res int) {
		if i == m {
			if loan == 0 && leading_zero == 3 {
				return 1
			}
			return 0
		}

		if dp[i][loan][leading_zero] != -1 {
			return dp[i][loan][leading_zero]
		}
		defer func() {
			dp[i][loan][leading_zero] = res
		}()
		// 如果loan = 1, 这里要有进位
		a := leading_zero & 1
		b := (leading_zero >> 1) & 1

		s1 := 1
		if a == 0 {
			s1 = 0
		}
		s2 := 1
		if b == 0 {
			s2 = 0
		}

		// 现在a/b都不能为0
		for c := range 2 {
			for x := s1; x <= 9; x++ {
				for y := s2; y <= 9; y++ {
					if (x+y+c)/10 == loan && (x+y+c)%10 == digits[i] {
						next := leading_zero
						if x != 0 {
							next |= 1
						}
						if y != 0 {
							next |= 2
						}

						res += f(i+1, c, next)
					}
				}
			}
		}
		return
	}

	res := f(0, 0, 0)

	return int64(res)
}

func bruteForce(n int64) int64 {
	var res int64
	for a := int64(1); a < n; a++ {
		if !hasZero(a) && !hasZero(n-a) {
			res++
		}
	}

	return res
}

func hasZero(n int64) bool {
	for n > 0 {
		if n%10 == 0 {
			return true
		}
		n /= 10
	}
	return false
}
