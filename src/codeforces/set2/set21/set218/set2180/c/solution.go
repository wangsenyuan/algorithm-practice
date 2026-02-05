package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		res := solve(n, k)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func solve(n int, k int) []int {
	res := make([]int, k)
	for i := range k {
		res[i] = n
	}
	if k&1 == 1 {
		return res
	}
	var free int
	for d := 30; d >= 0; d-- {
		if (n>>d)&1 == 1 {
			i := min(k-1, free)
			res[i] ^= 1 << d
			free = min(free+1, k)
		} else {
			// 前w个数，把这一位设置成1
			w := free - free&1
			for i := range w {
				res[i] |= 1 << d
			}
		}
	}
	return res
}
