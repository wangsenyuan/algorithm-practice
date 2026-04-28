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
		var x, y, k int
		fmt.Fscan(reader, &x, &y, &k)
		res := solve(x, y, k)
		fmt.Fprintln(writer, res)
	}
}

func solve(x int, y int, k int) int {
	if y == 1 {
		return -1
	}
	for i := 0; i < x; {
		cur := (k - 1) / (y - 1)
		if cur == 0 {
			break
		}
		fk := (cur+1)*(y-1) + 1
		cnt := (fk - k + cur - 1) / cur
		cnt = min(x-i, cnt)
		k += cnt * cur
		if k > 1e12 {
			return -1
		}
		i += cnt
	}

	return k
}
