package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		ans, ok := drive(reader)
		if !ok {
			break
		}
		fmt.Println(ans)
	}
}

func drive(reader *bufio.Reader) (int, bool) {
	var n, w int
	if _, err := fmt.Fscan(reader, &n, &w); err != nil || n < 0 {
		return 0, false
	}
	items := make([][2]int, n)
	for i := range items {
		fmt.Fscan(reader, &items[i][0], &items[i][1])
	}
	return solve(w, items), true
}

func solve(W int, items [][2]int) int {
	// n := len(items)
	width := bits.Len(uint(W))
	const mx = 1000
	f := make([][mx + 1]int, width)
	for _, cur := range items {
		w, v := cur[0], cur[1]
		tz := bits.TrailingZeros(uint(w))
		if tz >= width {
			continue
		}
		w >>= tz
		for j := mx; j >= w; j-- {
			f[tz][j] = max(f[tz][j], f[tz][j-w]+v)
		}
	}

	for i := 1; i < width; i++ {
		for j := mx; j >= 0; j-- {
			for k := range j + 1 {
				mx1 := min(k<<1|W>>(i-1)&1, mx)
				f[i][j] = max(f[i][j], f[i][j-k]+f[i-1][mx1])
			}
		}
	}
	return f[width-1][1]
}
