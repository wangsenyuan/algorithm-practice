package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	a := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func solve(a []int) int {
	// len(a) = 4
	d := a[1] - a[0]

	checkArithmetic := func(d int) bool {
		for i := 2; i < len(a); i++ {
			if a[i]-a[i-1] != d {
				return false
			}
		}
		return true
	}

	if checkArithmetic(d) {
		return a[3] + d
	}

	if a[0] == a[1] {
		return 42
	}
	c := gcd(a[0], a[1])
	p := pair{a[1] / c, a[0] / c}

	for i := 1; i < len(a); i++ {
		// a[i] / a[i-1] = a[1] / a[0]
		w := gcd(a[i], a[i-1])
		q := pair{a[i] / w, a[i-1] / w}
		if p != q {
			return 42
		}
	}
	res := a[3] * p.first / p.second
	if res*p.second != a[3]*p.first {
		return 42
	}

	return res
}
