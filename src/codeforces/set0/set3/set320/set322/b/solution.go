package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var r, g, b int
	fmt.Fscan(reader, &r, &g, &b)
	res := solve(r, g, b)
	fmt.Println(res)
}
func solve(r, g, b int) int {
	// 3 mixing bouquets = 1 of each color bouquet, so optimal w in {0,1,2}
	best := 0
	for w := range 3 {
		if w > min(r, g, b) {
			break
		}
		best = max(best, (r-w)/3+(g-w)/3+(b-w)/3+w)
	}
	return best
}
