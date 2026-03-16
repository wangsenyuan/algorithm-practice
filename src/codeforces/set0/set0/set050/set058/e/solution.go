package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	expr := readString(reader)
	res := solve(expr)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

type data struct {
	i, j, k, carry      int
	stopX, stopY        int
	topX, topY          int
	u, v, w             int
}

type item struct {
	i, j, k              int
	carry                int
	stopX, stopY         int
	topX, topY           int
	priority             int
	index                int
}

type PriorityQueue []*item

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
	item := x.(*item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func parseExpr(expr string) (a string, b string, c string) {
	l := strings.Index(expr, "+")
	r := strings.Index(expr, "=")
	a = reverse(expr[:l])
	b = reverse(expr[l+1 : r])
	c = reverse(expr[r+1:])
	return
}
func solve(expr string) string {
	a, b, c := parseExpr(expr)

	n := len(a)
	m := len(b)
	w := len(c)

	total := (n + 1) * (m + 1) * (w + 1) * 2 * 2 * 2 * 11 * 11
	dp := make([]int, total)
	fp := make([]data, total)
	for i := range total {
		dp[i] = inf
		fp[i].i = -1
	}

	getID := func(i, j, k, carry, stopX, stopY, topX, topY int) int {
		res := i
		res = res*(m+1) + j
		res = res*(w+1) + k
		res = res*2 + carry
		res = res*2 + stopX
		res = res*2 + stopY
		res = res*11 + topX
		res = res*11 + topY
		return res
	}

	var pq PriorityQueue

	update := func(i int, j int, k int, carry int, stopX int, stopY int, topX int, topY int, val int) bool {
		id := getID(i, j, k, carry, stopX, stopY, topX, topY)
		if dp[id] <= val {
			return false
		}
		dp[id] = val
		heap.Push(&pq, &item{i, j, k, carry, stopX, stopY, topX, topY, val, 0})
		return true
	}

	update(0, 0, 0, 0, 0, 0, 10, 10, 0)

	isGoal := func(i int, j int, k int, carry int, stopX int, stopY int, topX int, topY int) bool {
		if i != n || j != m || k != w || carry != 0 {
			return false
		}
		okX := stopX == 1 || topX > 0
		okY := stopY == 1 || topY > 0
		return okX && okY
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*item)
		id := getID(it.i, it.j, it.k, it.carry, it.stopX, it.stopY, it.topX, it.topY)
		if it.priority > dp[id] {
			continue
		}

		if isGoal(it.i, it.j, it.k, it.carry, it.stopX, it.stopY, it.topX, it.topY) {
			break
		}

		var digitsX []int
		if it.stopX == 1 {
			digitsX = []int{-1}
		} else {
			if it.i == n && it.topX > 0 {
				digitsX = append(digitsX, -1)
			}
			for u := 0; u < 10; u++ {
				digitsX = append(digitsX, u)
			}
		}
		var digitsY []int
		if it.stopY == 1 {
			digitsY = []int{-1}
		} else {
			if it.j == m && it.topY > 0 {
				digitsY = append(digitsY, -1)
			}
			for v := 0; v < 10; v++ {
				digitsY = append(digitsY, v)
			}
		}

		for _, u := range digitsX {
			for _, v := range digitsY {
				if u == -1 && v == -1 && it.carry == 0 {
					continue
				}

				sum := it.carry
				if u >= 0 {
					sum += u
				}
				if v >= 0 {
					sum += v
				}
				c1 := sum / 10
				z := sum % 10
				var add int
				i, j, k := it.i, it.j, it.k
				stopX, stopY := it.stopX, it.stopY
				topX, topY := it.topX, it.topY

				if u == -1 {
					stopX = 1
				} else {
					topX = u
					if i < n && u == int(a[i]-'0') {
						i++
					} else {
						add++
					}
				}
				if v == -1 {
					stopY = 1
				} else {
					topY = v
					if j < m && v == int(b[j]-'0') {
						j++
					} else {
						add++
					}
				}
				if k < w && z == int(c[k]-'0') {
					k++
				} else {
					add++
				}
				if update(i, j, k, c1, stopX, stopY, topX, topY, it.priority+add) {
					fp[getID(i, j, k, c1, stopX, stopY, topX, topY)] = data{
						it.i, it.j, it.k, it.carry, it.stopX, it.stopY, it.topX, it.topY,
						u, v, z,
					}
				}
			}
		}
	}

	// 回溯构造答案
	i, j, k, carry := n, m, w, 0
	stopX, stopY, topX, topY := 0, 0, 10, 10
	best := inf
	for sx := 0; sx < 2; sx++ {
		for sy := 0; sy < 2; sy++ {
			for tx := 0; tx < 11; tx++ {
				for ty := 0; ty < 11; ty++ {
					if !isGoal(n, m, w, 0, sx, sy, tx, ty) {
						continue
					}
					id := getID(n, m, w, 0, sx, sy, tx, ty)
					if dp[id] < best {
						best = dp[id]
						stopX, stopY = sx, sy
						topX, topY = tx, ty
					}
				}
			}
		}
	}

	var x []byte
	var y []byte
	var z []byte

	for {
		tmp := fp[getID(i, j, k, carry, stopX, stopY, topX, topY)]
		if tmp.u >= 0 {
			x = append(x, byte(tmp.u+'0'))
		}
		if tmp.v >= 0 {
			y = append(y, byte(tmp.v+'0'))
		}
		if tmp.w >= 0 {
			z = append(z, byte(tmp.w+'0'))
		}
		i, j, k = tmp.i, tmp.j, tmp.k
		carry = tmp.carry
		stopX, stopY = tmp.stopX, tmp.stopY
		topX, topY = tmp.topX, tmp.topY
		if i == 0 && j == 0 && k == 0 && carry == 0 && stopX == 0 && stopY == 0 && topX == 10 && topY == 10 {
			break
		}
	}

	a = string(x)
	b = string(y)
	c = string(z)

	return fmt.Sprintf("%s+%s=%s", a, b, c)
}

const inf = 1 << 60

func reverse(s string) string {
	buf := []byte(s)
	slices.Reverse(buf)
	return string(buf)
}
