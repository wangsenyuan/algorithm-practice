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
	ask := func(w int) int {
		fmt.Println("?", w)
		var h int
		fmt.Fscan(reader, &h)
		return h
	}
	res := solve(n, ask)
	fmt.Println("!", res)
}

func solve(n int, ask func(w int) int) int {
	// n * 2000
	S := sort.Search(n*2000+n-1, func(w int) bool {
		return w > 0 && ask(w) == 1
	})
	// S = sum(L) + n - 1

	best := S

	for h := 2; h <= n; h++ {
		w := S / h
		if w == 0 {
			break
		}
		tmp := ask(w)
		if tmp == 0 {
			break
		}
		best = min(best, w*tmp)
	}

	return best
}
