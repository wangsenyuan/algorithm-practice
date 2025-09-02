package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) int {
	var W, H int
	fmt.Fscan(reader, &W, &H)
	var first [4]int
	for i := 0; i < 4; i++ {
		fmt.Fscan(reader, &first[i])
	}
	var second [2]int
	for i := 0; i < 2; i++ {
		fmt.Fscan(reader, &second[i])
	}
	return solve(W, H, first[:], second[:])
}

func solve(W int, H int, first []int, second []int) int {
	// 分别将second放置到四个角落，check first的移动距离
	// (0, 0), (0, h), (w, 0), (w, h)
	w, h := second[0], second[1]
	x1, y1, x2, y2 := first[0], first[1], first[2], first[3]
	if w <= x1 || h <= y1 || x2 <= W-w || y2 <= H-h {
		return 0
	}
	// x1 < w & y1 < h
	var ans int = 1e18
	if w+x2-x1 <= W {
		ans = min(ans, (w - x1))
		ans = min(ans, (x2 - (W - w)))
	}
	if h+y2-y1 <= H {
		ans = min(ans, (h - y1))
		ans = min(ans, (y2 - (H - h)))
	}
	if ans >= max(W, H) {
		return -1
	}
	return ans
}
