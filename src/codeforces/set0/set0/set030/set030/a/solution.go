package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var a, b, n int
	fmt.Fscan(reader, &a, &b, &n)
	return solve(a, b, n)
}

func solve(a, b, n int) string {
	// a * pow(x, n) = b
	if a == 0 {
		if b == 0 {
			return "1"
		}
		return "No solution"
	}
	if b == 0 {
		return "0"
	}
	if abs(b)%abs(a) != 0 {
		return "No solution"
	}
	if a*b < 0 && n%2 == 0 {
		return "No solution"
	}
	w := abs(b) / abs(a)
	// w = pow(x, n)

	for x := 1; ; x++ {
		y := pow(x, n, w)
		if abs(y) == w {
			if a*b < 0 {
				return fmt.Sprintf("%d", -x)
			}
			return fmt.Sprintf("%d", x)
		}
		if y > w {
			break
		}
	}

	return "No solution"
}

func pow(a int, b int, w int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			if res > w/a || res*a > w {
				return w + 1
			}
			res *= a
		}

		a *= a

		b >>= 1
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
