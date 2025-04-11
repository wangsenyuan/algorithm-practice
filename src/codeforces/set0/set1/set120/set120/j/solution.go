package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	res, _ := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %d %d %d\r\n", res[0], res[1], res[2], res[3]))

	fmt.Fprint(w, buf.String())
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

func process(reader *bufio.Reader) (res []int, vecs [][]int) {
	n := readNum(reader)
	vecs = make([][]int, n)
	for i := range n {
		vecs[i] = readNNums(reader, 2)
	}
	res = solve(vecs)
	return
}

type pair struct {
	x int
	y int
}

func dist(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}

const inf = 1 << 60

func solve(vectors [][]int) (ans []int) {
	// 全部转换到第一象限
	n := len(vectors)
	vecs := make([][]int, n)
	for i, cur := range vectors {
		x, y := cur[0], cur[1]
		vecs[i] = []int{abs(x), abs(y), i}
	}

	best := inf

	var res []int

	update := func(a, b []int) {
		dx := a[0] - b[0]
		dy := a[1] - b[1]
		tmp := dx*dx + dy*dy
		if tmp < best {
			best = tmp
			res = []int{a[2], b[2]}
		}
	}

	cmp_x := func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	}

	cmp_y := func(a, b []int) int {
		if a[1] != b[1] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	}

	slices.SortFunc(vecs, cmp_x)

	var dfs func(l int, r int)

	buf := make([][]int, n)

	merge := func(l int, mid int, r int) {
		for i, j, k := 0, l, mid+1; i < r-l+1; i++ {
			if k > r || j <= mid && cmp_y(vecs[j], vecs[k]) <= 0 {
				buf[i] = vecs[j]
				j++
			} else {
				buf[i] = vecs[k]
				k++
			}
		}
	}

	dfs = func(l int, r int) {
		if r-l+1 <= 3 {
			for i := l; i <= r; i++ {
				for j := i + 1; j <= r; j++ {
					update(vecs[i], vecs[j])
				}
			}
			slices.SortFunc(vecs[l:r+1], cmp_y)
			return
		}
		mid := (l + r) / 2
		x0 := vecs[mid][0]
		dfs(l, mid)
		dfs(mid+1, r)
		merge(l, mid, r)
		copy(vecs[l:r+1], buf[0:r-l+1])
		var tsz int
		for i := l; i <= r; i++ {
			x := vecs[i][0]
			if (x-x0)*(x-x0) < best {
				for j := tsz - 1; j >= 0 && dist(vecs[i], buf[j]) < best; j-- {
					update(vecs[i], buf[j])
				}
				buf[tsz] = vecs[i]
				tsz++
			}
		}
	}

	dfs(0, n-1)

	i := res[0]
	j := res[1]
	var k int
	if sign(vectors[i][0]) == sign(vectors[j][0]) {
		k |= 1
	}
	if sign(vectors[i][1]) == sign(vectors[j][1]) {
		k |= 2
	}
	k++

	return []int{i + 1, 1, j + 1, k}
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func abs(x int) int {
	return max(x, -x)
}

func bruteForce(vecs [][]int) []int {
	best := inf
	var res []int

	get := func(v []int, k int) []int {
		x, y := v[0], v[1]
		if k&1 == 1 {
			x *= -1
		}
		if k&2 == 2 {
			y *= -1
		}
		return []int{x, y}
	}

	for i := 0; i < len(vecs); i++ {
		for j := i + 1; j < len(vecs); j++ {
			for k1 := 0; k1 < 4; k1++ {
				for k2 := 0; k2 < 4; k2++ {
					v1 := get(vecs[i], k1)
					v2 := get(vecs[j], k2)
					x := v1[0] + v2[0]
					y := v1[1] + v2[1]
					tmp := x*x + y*y
					if tmp < best {
						best = tmp
						res = []int{i + 1, k1 + 1, j + 1, k2 + 1}
					}
				}
			}
		}
	}
	return res
}
