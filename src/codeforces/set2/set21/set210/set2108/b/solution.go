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
		var n, x int
		fmt.Fscan(reader, &n, &x)
		res := solve(n, x)
		fmt.Fprintln(writer, res)
	}
}

func solve(n int, x int) int {
	if x == 0 {
		if n%2 == 0 {
			// n个1
			return n
		}
		// 1111 2
		if n == 1 {
			return -1
		}
		return (n - 2) + 2 + 3
	}

	if x == 1 {
		if n&1 == 1 {
			return n
		}
		return n + 3
	}

	// d < n
	// 对最低的两位要特殊处理
	res := x
	n -= bits.OnesCount(uint(x))
	if n <= 0 {
		return res
	}

	if n%2 == 0 {
		res += n
	} else {
		if x == 1 {

		}
		res += n + 1
	}

	return res
}
