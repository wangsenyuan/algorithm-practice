package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	res := solve(n, k)
	fmt.Println(res)
}

func solve(n int, k int) int {
	if n == 1 {
		return 0
	}
	if n <= k {
		// 1, 2 ... k头，直接转换就可以了
		return 1
	}
	// n > k
	// 全部用上，最后产生的头
	// 1 + 2 - 1 + 3 - 1 + 。。。 + k - 1
	tot := 1 + (1+k-1)*(k-1)/2
	if tot < n {
		return -1
	}
	// 貌似肯定有答案
	// 尽量使用后面的
	n--
	check := func(i int) bool {
		// i.....k
		sum := (i - 1 + k - 1) * (k - i + 1) / 2
		return sum < n
	}

	i := sort.Search(k, check)
	i--

	return k - i + 1
}
