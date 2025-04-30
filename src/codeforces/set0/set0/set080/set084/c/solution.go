package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cnt, res := process(reader)
	fmt.Println(cnt)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (hit int, res []int) {
	n := readNum(reader)
	targets := make([][]int, n)
	for i := range n {
		targets[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	shoots := make([][]int, m)
	for i := range m {
		shoots[i] = readNNums(reader, 2)
	}
	return solve(targets, shoots)
}

type Target struct {
	x  int
	y  int
	r  int
	id int
}

type Shoot struct {
	x int
	y int
}

func check(t Target, s Shoot) bool {
	dx := s.x - t.x
	dy := s.y - t.y
	return dx*dx+dy*dy <= t.r*t.r
}

const inf = 1 << 60

func solve(targets [][]int, shoots [][]int) (hit int, ans []int) {
	ts := make([]Target, len(targets))
	for i, cur := range targets {
		ts[i] = Target{cur[0], 0, cur[1], i}
	}

	ss := make([]Shoot, len(shoots))
	for i, cur := range shoots {
		ss[i] = Shoot{cur[0], cur[1]}
	}

	slices.SortFunc(ts, func(a, b Target) int {
		return a.x - b.x
	})

	m := len(ts)

	ans = make([]int, m)
	for i := range m {
		ans[i] = inf
	}
	for i := range ss {
		r := sort.Search(m, func(j int) bool {
			return ts[j].x-ts[j].r > ss[i].x
		})
		l := sort.Search(m, func(j int) bool {
			return ts[j].x+ts[j].r >= ss[i].x
		})

		// 因为不重叠，且在一条水平线上,所以 r - l < 3,
		for j := l; j < r; j++ {
			if check(ts[j], ss[i]) {
				ans[ts[j].id] = min(ans[ts[j].id], i+1)
			}
		}
	}
	for i := range m {
		if ans[i] == inf {
			ans[i] = -1
		} else {
			hit++
		}
	}
	return
}
