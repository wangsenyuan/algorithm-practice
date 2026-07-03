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

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) int {
	cnt := make(map[int]int)

	for num := 1; num <= n; num++ {
		x := num
		key := 1
		for p := 2; p*p <= x; p++ {
			if x%p != 0 {
				continue
			}
			c := 0
			for x%p == 0 {
				x /= p
				c++
			}
			if c&1 == 1 {
				key *= p
			}
		}
		if x > 1 {
			key *= x
		}
		cnt[key]++
	}

	var res int
	for _, c := range cnt {
		res += c * c
	}
	return res
}
