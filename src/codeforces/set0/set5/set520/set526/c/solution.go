package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var c, hr, hb, wr, wb int
	fmt.Fscan(reader, &c, &hr, &hb, &wr, &wb)
	return solve(c, hr, hb, wr, wb)
}

func solve(c int, hr int, hb int, wr int, wb int) int {
	c1 := int(math.Sqrt(float64(c)))

	if wr < wb {
		wr, wb = wb, wr
		hr, hb = hb, hr
	}
	var best int
	// wr >= wb
	if wr >= c1 {
		for i := 0; i*wr <= c; i++ {
			j := (c - i*wr) / wb
			tmp := i*hr + j*hb
			best = max(best, tmp)
		}
		return best
	}

	if wr*hb > wb*hr {
		wr, wb = wb, wr
		hr, hb = hb, hr
	}
	// wr * hb <= wb * hr
	// hb/wb <= hr/wr
	// 单位重量蓝色的价值低于红色的价值
	// 蓝色的不会超过c1
	for i := 0; i*wb <= c && i <= c1; i++ {
		j := (c - i*wb) / wr
		tmp := i*hb + j*hr
		best = max(best, tmp)
	}
	return best
}
