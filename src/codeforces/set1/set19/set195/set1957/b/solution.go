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
	for range tc {
		_, res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (int, []int) {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	res := solve(n, k)
	return k, res
}

func solve(n int, k int) []int {
	res := make([]int, n)

	if n == 1 {
		res[0] = k
		return res
	}
	h := bits.Len(uint(k))
	if h == 1 {
		res[0] = k
	} else {
		res[0] = 1<<(h-1) - 1
	}
	k -= res[0]
	res[1] = k

	return res
}
