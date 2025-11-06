package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var l, r, k int
	fmt.Fscan(reader, &l, &r, &k)
	res := solve(l, r, k)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}

	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func solve(l int, r int, k int) []int {
	var res []int

	for x := 1; x <= r; x *= k {
		if x >= l {
			res = append(res, x)
		}
		if x > r/k || x == r/k && x*k > r {
			break
		}
	}
	return res
}
