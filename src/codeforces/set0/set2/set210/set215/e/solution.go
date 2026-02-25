package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var l, r int
	fmt.Fscan(reader, &l, &r)
	res := solve(l, r)
	fmt.Println(res)
}

func solve(l int, r int) int {
	return play(r) - play(l-1)
}

func play(num int) int {

	var dfs func(n int, x int) int

	dfs = func(n int, x int) int {
		if n == 1 {
			return 0
		}
		var ans int
		for i := 1; i <= n/2; i++ {
			if n%i == 0 {
				var num int
				for j := 0; j < n; j += i {
					num |= 1 << j
				}
				ans += x/num - (1<<(n-1))/num - dfs(i, x/num)
			}
		}
		return ans
	}

	if num < 3 {
		return 0
	}
	var ans int
	for i := 2; 1<<(i-1) <= num; i++ {
		ans += dfs(i, min(num, 1<<i-1))
	}
	return ans
}

func bruteForce(l int, r int) int {
	var res int

	check := func(num int) bool {
		h := bits.Len(uint(num))
		for i := 1; i < h; i++ {
			if h%i == 0 {
				w := num >> (h - i)
				var tmp int
				for range h / i {
					tmp = tmp<<i | w
				}
				if tmp == num {
					return true
				}
			}
		}
		return false
	}

	for x := l; x <= r; x++ {
		if check(x) {
			res++
		}
	}
	return res
}
