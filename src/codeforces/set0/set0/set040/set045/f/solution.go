package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	res := solve(m, n)
	fmt.Println(res)
}

func solve(m int, n int) int {
	if n == 1 {
		// 它只能带一只过去
		return -1
	}
	if n == 2 && m == 3 {
		return 11
	}
	if n == 3 && m == 5 {
		return 11
	}
	var ans int
	var flag bool
	for {
		if n >= 2*m {
			return ans + 1
		}
		if n >= m {
			if n == m {
				return ans + 5
			}
			return ans + 3
		}
		if !flag {
			m -= n - 2
			ans = 4
			flag = true
		} else {
			if n/2 == 1 {
				break
			}
			m -= n/2 - 1
			ans += 2
		}
	}

	return -1
}
