package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	h := sort.Search(min(n, 10000000), func(i int) bool {
		// 共有i层, 每层有 3 * i - 1 张牌
		sum := (1 + i) * i / 2
		return 3*sum-i > n
	})
	h--

	for (h+n)%3 != 0 {
		h--
	}

	return (h + 2) / 3
}
