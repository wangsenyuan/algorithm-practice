package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var k int
	fmt.Fscan(reader, &k)
	res := solve(k)
	fmt.Println(res)
}

func solve(k int) int {
	check := func(x int, y int) bool {
		if (x+1)*(x+1)+3*(y+1)*(y+1) > 4*k*k {
			return false
		}
		if (x+2)*(x+2)+3*y*y > 4*k*k {
			return false
		}
		return true
	}

	var x, y int
	for check(x, y) {
		y += 2
	}
	ans := -y - 3

	for y >= 0 {
		ans += 2 * (y + 1)
		x += 3
		y += 1
		for y >= 0 && !check(x, y) {
			y -= 2
		}
	}
	return ans
}
