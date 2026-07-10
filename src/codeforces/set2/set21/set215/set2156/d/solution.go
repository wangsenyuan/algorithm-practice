package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)

	ask := func(i, x int) int {
		fmt.Fprintf(writer, "? %d %d\n", i, x)
		writer.Flush()
		var res int
		fmt.Fscan(reader, &res)
		return res
	}

	for range tc {
		var n int
		fmt.Fscan(reader, &n)
		ans := solve(n, ask)
		fmt.Fprintf(writer, "! %d\n", ans)
		writer.Flush()
	}
}

func solve(n int, ask func(i, x int) int) int {
	h := bits.Len(uint(n))

	arr := make([]int, n-1)
	for i := range n - 1 {
		arr[i] = i + 1
	}

	var res int
	for d := range h {
		x := 1 << d
		var arr0 []int
		var arr1 []int
		for _, i := range arr {
			y := ask(i, x)
			if y == 0 {
				arr0 = append(arr0, i)
			} else {
				arr1 = append(arr1, i)
			}
		}

		mod := 1 << (d + 1)
		cnt0 := count(n, mod, res)
		if len(arr0) < cnt0 {
			arr = arr0
		} else {
			res |= x
			arr = arr1
		}
	}

	return res
}

func count(n int, mod int, rem int) int {
	// rem, rem + mod, rem + 2 * mod, ...
	first := rem
	if first == 0 {
		first = mod
	}
	if first > n {
		return 0
	}
	return (n-first)/mod + 1
}
