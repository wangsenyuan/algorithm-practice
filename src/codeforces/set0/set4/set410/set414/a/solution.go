package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)
	res := solve(n, k)
	if len(res) == 0 {
		fmt.Println(-1)
	} else {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func solve(n int, k int) []int {
	if k < n/2 {
		return nil
	}
	if k == 0 {
		if n == 1 {
			return []int{1}
		}
		return nil
	}
	if n == 1 {
		return nil
	}

	var res []int

	x := k - (n-2)/2
	if x > 1 {
		if x*103 <= 1e9 {
			res = append(res, x*102, x*103)
		} else {
			res = append(res, x, 2*x)
		}
	} else {
		res = append(res, n, n-1)
	}

	for i := 2; i < n; i++ {
		res = append(res, n-i)
	}

	return res
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
