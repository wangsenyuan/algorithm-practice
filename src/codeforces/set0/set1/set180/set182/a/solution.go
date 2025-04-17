package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.6f\n", res)
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

func process(reader *bufio.Reader) float64 {
	a, b := readTwoNums(reader)
	line := readNNums(reader, 4)
	n := readNum(reader)
	lines := make([][]int, n)
	for i := range n {
		lines[i] = readNNums(reader, 4)
	}
	return solve(a, b, line[:2], line[2:], lines)
}

type Edge struct {
	to     int
	weight float64
}

const eps = 1e-6

const inf = 1 << 50

func checkDirection(line []int) int {
	if line[0] == line[2] {
		// 垂直线
		return 0
	}
	return 1
}

func distance(a []int, b []int) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func find1(a, b []int) float64 {
	// a是水平线，b是垂直线
	if a[0] <= b[0] && b[0] <= a[2] {
		d1 := abs(b[1] - a[1])
		d2 := abs(b[3] - a[1])
		return float64(min(d1, d2))
	}
	if b[1] <= a[1] && a[1] <= b[3] {
		d1 := abs(a[0] - b[0])
		d2 := abs(a[2] - b[0])
		return float64(min(d1, d2))
	}
	return find4(a, b)
}

func abs(num int) int {
	return max(num, -num)
}

func find2(a, b []int) float64 {
	// 两条垂直线，
	if a[3] < b[1] || b[3] < a[1] {
		// 没有水平重叠的区域
		return find4(a, b)
	}
	// 水平方向的最短距离
	return float64(abs(a[0] - b[0]))
}

func find3(a, b []int) float64 {
	// 两条水平线
	if a[2] < b[0] || b[2] < a[0] {
		// 没有水平重叠的区域
		return find4(a, b)
	}

	return float64(abs(a[1] - b[1]))
}

func find4(a, b []int) float64 {
	var res float64 = inf
	for i := range []int{0, 1} {
		for j := range []int{0, 1} {
			tmp := distance(a[2*i:], b[2*j:])
			res = min(res, tmp)
		}
	}
	return res
}

func findShortedDistance(a, b []int) float64 {
	da := checkDirection(a)
	db := checkDirection(b)
	if da != db {
		if da == 1 {
			return find1(a, b)
		}
		return find1(b, a)
	}
	if da == 0 {
		return find2(a, b)
	}
	return find3(a, b)
}

func solve(a int, b int, A []int, B []int, lines [][]int) float64 {
	n := len(lines)

	for i := range n {
		if lines[i][0] > lines[i][2] {
			lines[i][0], lines[i][2] = lines[i][2], lines[i][0]
		}
		if lines[i][1] > lines[i][3] {
			lines[i][1], lines[i][3] = lines[i][3], lines[i][1]
		}
	}

	dist := make([][]float64, n+2)

	for i := 0; i < n+2; i++ {
		dist[i] = make([]float64, n+2)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist[i][j] = findShortedDistance(lines[i], lines[j])
			dist[j][i] = dist[i][j]
		}
	}

	get := func(cur []int, line []int) float64 {
		res := min(distance(cur, line), distance(cur, line[2:]))
		di := checkDirection(line)
		if di == 0 {
			if line[1] <= cur[1] && cur[1] <= line[3] {
				res = min(res, float64(abs(cur[0]-line[0])))
			}
		} else {
			if line[0] <= cur[0] && cur[0] <= line[2] {
				res = min(res, float64(abs(cur[1]-line[1])))
			}
		}
		return res
	}

	for i := 0; i < n; i++ {
		dist[i][n] = get(A, lines[i])
		dist[n][i] = dist[i][n]
		dist[i][n+1] = get(B, lines[i])
		dist[n+1][i] = dist[i][n+1]
	}

	if distance(A, B) <= float64(a) {
		dist[n][n+1] = distance(A, B)
		dist[n+1][n] = dist[n][n+1]
	} else {
		dist[n][n+1] = inf
		dist[n+1][n] = inf
	}

	items := make([]*Item, n+2)
	pq := make(PriorityQueue, n+2)
	for i := range n + 2 {
		it := new(Item)
		it.id = i
		it.priority = inf
		it.index = i
		items[i] = it
		pq[i] = it
	}

	items[n].priority = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority >= inf {
			break
		}
		u := it.id
		if u == n+1 {
			return it.priority
		}
		t := (int(it.priority) + a + b - 1) / (a + b) * (a + b)
		// t是离开的最早的时间
		for i := range n + 2 {
			if items[i].index >= 0 && items[i].priority > float64(t)+dist[u][i] && dist[u][i] <= float64(a) {
				pq.update(items[i], float64(t)+dist[u][i])
			}
		}
	}

	return -1
}

type Item struct {
	id       int
	priority float64
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

func (pq *PriorityQueue) update(it *Item, v float64) {
	it.priority = v
	heap.Fix(pq, it.index)
}
