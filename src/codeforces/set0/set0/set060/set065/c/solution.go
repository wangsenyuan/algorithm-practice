package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ok, ans, p0 := process(reader)
	if !ok {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Printf("%.10f\n", ans)
		fmt.Printf("%.10f %.10f %.10f\n", p0[0], p0[1], p0[2])
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func process(reader *bufio.Reader) (bool, float64, []float64) {
	n := readNum(reader)
	coordinates := make([][]int, n+1)
	for i := range n + 1 {
		coordinates[i] = readNNums(reader, 3)
	}
	v := readNNums(reader, 2)
	p := readNNums(reader, 3)

	return solve(coordinates, v, p)
}

const eps = 1e-12

func solve(coordinates [][]int, v []int, p []int) (bool, float64, []float64) {
	n := len(coordinates)
	vp, vs := float64(v[0]), float64(v[1])

	harry_pos := convertToFloat(p)

	get_position := func(i int, t, t0, t1 float64) []float64 {
		cur := coordinates[i]
		next := coordinates[i+1]

		dt := t - t0
		x := float64(cur[0]) + dt/(t1-t0)*(float64(next[0])-float64(cur[0]))
		y := float64(cur[1]) + dt/(t1-t0)*(float64(next[1])-float64(cur[1]))
		z := float64(cur[2]) + dt/(t1-t0)*(float64(next[2])-float64(cur[2]))

		return []float64{x, y, z}
	}

	check := func(t float64, t0 float64, t1 float64, i int) bool {
		// 是否可以在时刻t内，harry到到节点
		pos := get_position(i, t, t0, t1)
		t2 := travel(harry_pos, pos) / vp
		return t2 <= t+eps
	}

	var t0 float64

	for i := 0; i+1 < n; i++ {
		t1 := travel(convertToFloat(coordinates[i]), convertToFloat(coordinates[i+1]))
		t1 /= vs

		lo, hi := 0.0, t1

		res := -1.0

		for range 100 {
			mid := (lo + hi) / 2
			if check(t0+mid, t0, t0+t1, i) {
				res = t0 + mid
				hi = mid
			} else {
				lo = mid
			}
		}

		if res >= t0-eps {
			return true, res, get_position(i, res, t0, t0+t1)
		}

		t0 += t1
	}

	return false, 0, nil
}

func convertToFloat(a []int) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = float64(x)
	}
	return res
}

func travel(a, b []float64) float64 {
	dx := b[0] - a[0]
	dy := b[1] - a[1]
	dz := b[2] - a[2]

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
