package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	buf := make([]byte, 4096)
	var _i, _n int
	rc := func() byte {
		if _i == _n {
			_n, _ = reader.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rd := func() (x int) {
		b := rc()
		for ; '0' > b && b != '-'; b = rc() {
		}
		sign := 1
		if b == '-' {
			sign = -1
			b = rc()
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		x *= sign
		return
	}

	n := rd()
	m := rd()
	trees := make([][]int, n)
	for i := range n {
		trees[i] = []int{rd(), rd(), rd(), rd()}
	}
	mushrooms := make([][]int, m)
	for i := range m {
		mushrooms[i] = []int{rd(), rd()}
	}
	return solve(trees, mushrooms)
}

type BIT []float64

func (bit BIT) update(i int, v float64) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) get(i int) float64 {
	i++
	var res float64
	for i > 0 {

		res += bit[i]
		i -= i & -i
	}
	return res
}

func solve(trees [][]int, mushrooms [][]int) float64 {
	var pos []int
	for _, tree := range trees {
		x, h := tree[0], tree[1]
		pos = append(pos, x-h, x, x+1, x+h+1)
	}
	for _, cur := range mushrooms {
		pos = append(pos, cur[0])
	}

	slices.Sort(pos)
	pos = slices.Compact(pos)

	m := len(pos)

	diff := make(BIT, m+3)

	flag := make(BIT, m+3)

	for _, tree := range trees {
		x, h, l, r := tree[0], tree[1], tree[2], tree[3]
		i1 := sort.SearchInts(pos, x-h)
		i2 := sort.SearchInts(pos, x)
		i3 := sort.SearchInts(pos, x+1)
		i4 := sort.SearchInts(pos, x+h+1)

		// -inf will cause problem
		p1 := math.Log(float64(100-l) / 100.0)
		p2 := math.Log(float64(100-r) / 100.0)

		if l < 100 && r < 100 {
			diff.update(i1, p1)
			diff.update(i2, -p1)
			diff.update(i3, p2)
			diff.update(i4, -p2)
		} else if l == 100 {
			// 左边区域肯定全部被覆盖了
			flag.update(i1, 1)
			flag.update(i2, -1)
		} else {
			flag.update(i3, 1)
			flag.update(i4, -1)
		}
	}
	var ans float64

	for _, cur := range mushrooms {
		x, z := cur[0], cur[1]
		i := sort.SearchInts(pos, x)

		if flag.get(i) > 0 {
			continue
		}

		y := diff.get(i)
		ans += float64(z) * math.Exp(y)
	}

	return ans
}

func solve1(trees [][]int, mushrooms [][]int) float64 {
	var pos []int
	for _, tree := range trees {
		x, h := tree[0], tree[1]
		pos = append(pos, x-h, x, x+1, x+h+1)
	}
	for _, cur := range mushrooms {
		pos = append(pos, cur[0])
	}

	slices.Sort(pos)
	pos = slices.Compact(pos)

	m := len(pos)

	diff := make([][101]int16, m+1)

	for _, tree := range trees {
		x, h, l, r := tree[0], tree[1], tree[2], tree[3]
		i1 := sort.SearchInts(pos, x-h)
		i2 := sort.SearchInts(pos, x)
		i3 := sort.SearchInts(pos, x+1)
		i4 := sort.SearchInts(pos, x+h+1)
		diff[i1][l]++
		diff[i2][l]--
		diff[i3][r]++
		diff[i4][r]--
	}

	at := make([][]int, m)
	for i, cur := range mushrooms {
		i1 := sort.SearchInts(pos, cur[0])
		at[i1] = append(at[i1], i)
	}

	var ans float64

	for i := range m {
		if i > 0 {
			for j := range 101 {
				diff[i][j] += diff[i-1][j]
			}
		}
		// log(survive) = sum of c*log(d) to avoid underflow
		logSurvive := 0.0
		for j, c := range diff[i] {
			if c == 0 {
				continue
			}
			if j == 100 {
				logSurvive = math.Inf(-1)
				break
			}
			d := float64(100-j) / 100.0
			logSurvive += float64(c) * math.Log(d)
		}
		survive := math.Exp(logSurvive) // 0 when logSurvive == -Inf

		for _, j := range at[i] {
			ans += float64(mushrooms[j][1]) * survive
		}
	}

	return ans
}
