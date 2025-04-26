package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	cables := make([][]int, m)
	for i := 0; i < m; i++ {
		cables[i] = readNNums(reader, 2)
	}
	q := readNum(reader)
	experiments := make([][]int, q)
	for i := 0; i < q; i++ {
		experiments[i] = readNNums(reader, 2)
	}
	return solve(n, cables, experiments)
}

const inf = 1 << 30

func solve(n int, cables [][]int, experiments [][]int) []int {
	m := len(cables)
	pref := make([]*DSU, m+1)
	pref[0] = NewDSU(n)
	for i := 0; i < m; i++ {
		pref[i+1] = pref[i].Copy()
		pref[i+1].Union(cables[i][0]-1, cables[i][1]-1)
	}

	suf := make([]*DSU, m+1)
	suf[m] = NewDSU(n + 1)
	for i := m - 1; i >= 0; i-- {
		suf[i] = suf[i+1].Copy()
		suf[i].Union(cables[i][0]-1, cables[i][1]-1)
	}

	ans := make([]int, len(experiments))
	tmp := NewDSU(n + 1)
	for i, cur := range experiments {
		l, r := cur[0], cur[1]
		l--
		r--
		copy(tmp.arr, pref[l].arr)
		copy(tmp.cnt, pref[l].cnt)
		tmp.size = pref[l].size
		for u := 0; u < n; u++ {
			v := suf[r+1].Find(u)
			tmp.Union(u, v)
		}

		ans[i] = tmp.size
	}

	return ans
}

type DSU struct {
	arr  []int
	cnt  []int
	size int
}

func NewDSU(n int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = n
	return set
}

func (set *DSU) Find(u int) int {
	if set.arr[u] != u {
		set.arr[u] = set.Find(set.arr[u])
	}
	return set.arr[u]
}

func (set *DSU) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	set.size--
	return true
}

func (set *DSU) Reset() {
	for i := range set.arr {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = len(set.arr)
}

func (set *DSU) Copy() *DSU {
	newSet := NewDSU(len(set.arr))
	copy(newSet.arr, set.arr)
	copy(newSet.cnt, set.cnt)
	newSet.size = set.size
	return newSet
}
