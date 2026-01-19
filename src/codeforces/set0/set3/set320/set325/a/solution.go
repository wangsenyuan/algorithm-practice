package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	rects := make([][]int, n)
	for i := 0; i < n; i++ {
		rects[i] = make([]int, 4)
		fmt.Fscan(reader, &rects[i][0], &rects[i][1], &rects[i][2], &rects[i][3])
	}
	return solve(rects)
}

const inf = 1 << 60

func solve(rects [][]int) bool {
	var sum int
	x1, y1 := inf, inf
	x2, y2 := -inf, -inf

	for _, cur := range rects {
		x1 = min(x1, cur[0])
		y1 = min(y1, cur[1])
		x2 = max(x2, cur[2])
		y2 = max(y2, cur[3])
		dx := cur[2] - cur[0]
		dy := cur[3] - cur[1]
		sum += dx * dy
	}

	return sum == (x2-x1)*(y2-y1) && (x2-x1) == (y2-y1)
}
