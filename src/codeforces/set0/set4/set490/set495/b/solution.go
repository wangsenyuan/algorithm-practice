package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscan(reader, &a, &b)
	res := solve(a, b)
	if res < 0 {
		fmt.Println("infinity")
	} else {
		fmt.Println(res)
	}
}

func solve(a int, b int) int {
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	diff := a - b

	var res int

	add := func(x int) {
		if x > b {
			res++
		}
	}

	for x := 1; x <= diff/x; x++ {
		if diff%x == 0 {
			add(x)
			if x*x != diff {
				add(diff / x)
			}
		}
	}
	return res
}
