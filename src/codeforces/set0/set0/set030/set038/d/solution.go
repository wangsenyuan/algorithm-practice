package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	bricks := make([][]int, n)
	for i := range n {
		bricks[i] = readNNums(reader, 4)
	}
	return solve(bricks)
}

const eps = 1e-9

func checkLess(a, b float64) bool {
	return a < b-eps
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}

func solve(bricks [][]int) int {
	n := len(bricks)

	var f func(weight int, center []float64, i int) bool

	get := func(x0 float64, w0 float64, x1 float64, w1 float64) float64 {
		if checkLess(x1, x0) {
			x0, x1 = x1, x0
			w0, w1 = w1, w0
		}
		// x0 <= x1
		if math.Abs(x1-x0) < eps {
			return x0
		}
		// x0 < x1
		// nx更靠近重的那头
		// (nx - x0) / (x1 - nx) = w1 / w0
		// (nx - x0) * w0 = (x1 - nx) * w1
		// nx * (w0 + w1) = x0 * w0 + x1 * w1
		// nx = (x0 * w0 + x1 * w1) / (w0 + w1)
		return (x0*w0 + x1*w1) / (w0 + w1)
	}

	f = func(weight int, center []float64, i int) bool {
		if i < 0 {
			return true
		}
		cur := bricks[i]
		x0, y0 := center[0], center[1]
		x1, y1 := float64(cur[0]), float64(cur[1])
		x2, y2 := float64(cur[2]), float64(cur[3])
		if checkLess(x0, x1) || checkLess(y0, y1) || checkLess(x2, x0) || checkLess(y2, y0) {
			return false
		}
		// 重心会不会变成分数呢？
		w0 := float64(weight)
		w := pow(cur[2]-cur[0], 3)
		w1 := float64(w)
		x3, y3 := (x1+x2)/2, (y1+y2)/2

		nx := get(x0, w0, x3, w1)
		ny := get(y0, w0, y3, w1)
		return f(weight+w, []float64{nx, ny}, i-1)
	}

	check := func(m int) bool {
		cur := bricks[m]
		w := pow(cur[2]-cur[0], 3)
		return f(w, []float64{float64(cur[0]+cur[2]) / 2, float64(cur[1]+cur[3]) / 2}, m-1)
	}

	for i := 0; i < n; i++ {
		cur := bricks[i]
		if cur[0] > cur[2] {
			cur[0], cur[2] = cur[2], cur[0]
		}
		if cur[1] > cur[3] {
			cur[1], cur[3] = cur[3], cur[1]
		}
		if i > 0 && !check(i) {
			return i
		}
	}

	return n
}
