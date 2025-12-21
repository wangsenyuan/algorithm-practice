package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var m, n int
	fmt.Fscan(reader, &m, &n)
	res := solve(m, n)
	fmt.Printf("%.12f\n", res)
}

func solve(m int, n int) float64 {
	var res float64

	for i := 1; i <= m; i++ {
		// 如果i是最大值
		x := float64(i) / float64(m)
		x = math.Pow(x, float64(n))
		y := float64(i-1) / float64(m)
		y = math.Pow(y, float64(n))
		res += float64(i) * (x - y)
	}
	return res
}
