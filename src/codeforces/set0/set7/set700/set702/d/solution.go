package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var d, k, a, b, t int
	fmt.Fscan(reader, &d, &k, &a, &b, &t)
	return solve(d, k, a, b, t)
}

func solve(d int, k int, a int, b int, t int) int {
	if k >= d {
		return a * d
	}
	// k < d
	// 完全步行
	res := b * d
	w, v := d/k, d%k
	// w > 0, 完全开车
	res = min(res, a*d+w*t)
	// 只有第一段开车，然后步行
	res = min(res, a*k+(d-k)*b)
	// 最后一段步行
	res = min(res, a*k*w+(w-1)*t+v*b)

	return res
}
