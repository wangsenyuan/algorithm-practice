package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func drive(reader *bufio.Reader) int {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	initial := readNNums(reader, 2)
	hide := readNNums(reader, 2)
	return solve(a, initial, hide)
}

const inf = 1 << 60

var dd = []int{-1, 0, 1, 0, -1}

func solve(a []string, initial []int, hide []int) int {
	n := len(a)
	m := len(a[0])
	dist := make([][][]int, n)
	for i := range n {
		dist[i] = make([][]int, m)
		for j := range m {
			dist[i][j] = make([]int, 4)
			for k := range 4 {
				dist[i][j][k] = inf
			}
		}
	}

	var que [][]int
	push := func(x int, y int, z int, d int) {
		if dist[x][y][z] == inf {
			dist[x][y][z] = d
			que = append(que, []int{x, y, z, d})
		}
	}

	pop := func() (r int, c int, w int, d int) {
		r, c, w, d = que[0][0], que[0][1], que[0][2], que[0][3]
		que = que[1:]
		return
	}

	push(initial[0]-1, initial[1]-1, 0, 0)

	checkOpen := func(r int, c int, w int, i int) bool {
		// i = 0(up), 1(right), 2(down), 3(left)
		if a[r][c] == '+' {
			return true
		}
		if a[r][c] == '*' {
			return false
		}

		if a[r][c] == '-' {
			return i&1 != w&1
		}
		if a[r][c] == '|' {
			return i&1 == w&1
		}

		if a[r][c] == '^' {
			return i == w
		}
		if a[r][c] == '>' {
			// 门一开始是往右开的，当要往i走的时候，门必须转过来才行
			// w = 0的时候， i= 1， w = 1的时候, i = 2
			return i == (w+1)%4
		}
		if a[r][c] == 'v' {
			return i == (w+2)%4
		}
		if a[r][c] == '<' {
			return i == (w+3)%4
		}

		if a[r][c] == 'U' {
			// 向上没有门
			return i != w
		}

		if a[r][c] == 'R' {
			return i != (w+1)%4
		}

		if a[r][c] == 'D' {
			return i != (w+2)%4
		}

		return i != (w+3)%4
	}

	dp := make([][]int, n)
	for r := range n {
		dp[r] = make([]int, m)
		for c := range m {
			for i := range 4 {
				for w := range 4 {
					if checkOpen(r, c, w, i) {
						dp[r][c] |= (1 << (i*4 + w))
					}
				}
			}
		}
	}

	canPass := func(i int, r int, c int, nr int, nc int, w int) bool {
		if (dp[r][c]>>(i*4+w))&1 == 0 {
			return false
		}

		ni := (i + 2) % 4

		return (dp[nr][nc]>>(ni*4+w))&1 == 1
	}

	for len(que) > 0 {
		r, c, w, d := pop()
		if d == inf {
			break
		}
		if r == hide[0]-1 && c == hide[1]-1 {
			return d
		}
		// 看看能不能直接通过
		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			tmp := d + 1
			if nr >= 0 && nr < n && nc >= 0 && nc < m && canPass(i, r, c, nr, nc, w) {
				push(nr, nc, w, tmp)
			}
		}
		// push
		push(r, c, (w+1)%4, d+1)
	}

	return -1
}

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	it.index = -1
	return it
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
