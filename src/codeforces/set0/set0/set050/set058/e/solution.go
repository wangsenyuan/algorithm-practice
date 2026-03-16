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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return solve(strings.TrimSpace(s))
}

// parseExpr splits "a+b=c" and returns reversed a, b, c (LSB first).
func parseExpr(expr string) (string, string, string) {
	plus := strings.IndexByte(expr, '+')
	eq := strings.IndexByte(expr, '=')
	return rev(expr[:plus]), rev(expr[plus+1 : eq]), rev(expr[eq+1:])
}

func rev(s string) string {
	b := []byte(s)
	slices.Reverse(b)
	return string(b)
}

const (
	inf     = 1 << 60
	noDigit = 10 // sentinel: no digit emitted yet for this operand
)

// Dijkstra state for digit-by-digit construction of x+y=z from LSB to MSB.
type state struct {
	ia, ib, ic   int // positions matched in reversed a, b, c
	carry        int // addition carry (0 or 1)
	doneX, doneY int // 1 if operand x/y has finished emitting digits
	msdX, msdY   int // most recent (most significant so far) digit of x/y; noDigit if none
}

// transition records parent state and chosen digits for backtracking.
type transition struct {
	from       state
	dx, dy, dz int // digits appended to x, y, z (-1 = operand stopped)
}

func solve(expr string) string {
	a, b, c := parseExpr(expr)
	n, m, w := len(a), len(b), len(c)

	dims := [8]int{n + 1, m + 1, w + 1, 2, 2, 2, 11, 11}
	total := 1
	for _, d := range dims {
		total *= d
	}

	encode := func(s state) int {
		id := s.ia
		id = id*dims[1] + s.ib
		id = id*dims[2] + s.ic
		id = id*dims[3] + s.carry
		id = id*dims[4] + s.doneX
		id = id*dims[5] + s.doneY
		id = id*dims[6] + s.msdX
		id = id*dims[7] + s.msdY
		return id
	}

	dist := make([]int, total)
	parent := make([]transition, total)
	for i := range dist {
		dist[i] = inf
	}

	var pq priQueue
	relax := func(s state, cost int) bool {
		id := encode(s)
		if dist[id] <= cost {
			return false
		}
		dist[id] = cost
		heap.Push(&pq, &pqItem{s: s, cost: cost})
		return true
	}

	start := state{msdX: noDigit, msdY: noDigit}
	relax(start, 0)

	isGoal := func(s state) bool {
		if s.ia != n || s.ib != m || s.ic != w || s.carry != 0 {
			return false
		}
		// No leading zeros: top digit must be >0, or the operand was stopped (length determined by the other).
		okX := s.doneX == 1 || s.msdX > 0
		okY := s.doneY == 1 || s.msdY > 0
		return okX && okY
	}

	// digitChoices returns possible digits for an operand.
	// -1 means "stop emitting digits for this operand".
	digitChoices := func(done int, allMatched bool, msd int) []int {
		if done == 1 {
			return []int{-1}
		}
		var ds []int
		if allMatched && msd > 0 {
			ds = append(ds, -1)
		}
		for d := 0; d <= 9; d++ {
			ds = append(ds, d)
		}
		return ds
	}

	var goal state
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*pqItem)
		s := cur.s
		if cur.cost > dist[encode(s)] {
			continue
		}
		if isGoal(s) {
			goal = s
			break
		}

		for _, dx := range digitChoices(s.doneX, s.ia == n, s.msdX) {
			for _, dy := range digitChoices(s.doneY, s.ib == m, s.msdY) {
				if dx == -1 && dy == -1 && s.carry == 0 {
					continue
				}

				sum := s.carry + max(dx, 0) + max(dy, 0)
				dz := sum % 10

				ns := state{
					ia: s.ia, ib: s.ib, ic: s.ic,
					carry: sum / 10,
					doneX: s.doneX, doneY: s.doneY,
					msdX: s.msdX, msdY: s.msdY,
				}
				extra := 0

				if dx == -1 {
					ns.doneX = 1
				} else {
					ns.msdX = dx
					if s.ia < n && dx == int(a[s.ia]-'0') {
						ns.ia++
					} else {
						extra++
					}
				}
				if dy == -1 {
					ns.doneY = 1
				} else {
					ns.msdY = dy
					if s.ib < m && dy == int(b[s.ib]-'0') {
						ns.ib++
					} else {
						extra++
					}
				}
				if s.ic < w && dz == int(c[s.ic]-'0') {
					ns.ic++
				} else {
					extra++
				}

				if relax(ns, cur.cost+extra) {
					parent[encode(ns)] = transition{from: s, dx: dx, dy: dy, dz: dz}
				}
			}
		}
	}

	// Backtrack from goal to start to reconstruct x, y, z (digits collected MSB-first).
	var xb, yb, zb []byte
	for s := goal; s != start; {
		tr := parent[encode(s)]
		if tr.dx >= 0 {
			xb = append(xb, byte(tr.dx+'0'))
		}
		if tr.dy >= 0 {
			yb = append(yb, byte(tr.dy+'0'))
		}
		zb = append(zb, byte(tr.dz+'0'))
		s = tr.from
	}
	return fmt.Sprintf("%s+%s=%s", string(xb), string(yb), string(zb))
}

// --- min-heap priority queue ---

type pqItem struct {
	s    state
	cost int
	idx  int
}

type priQueue []*pqItem

func (q priQueue) Len() int           { return len(q) }
func (q priQueue) Less(i, j int) bool { return q[i].cost < q[j].cost }
func (q priQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i]; q[i].idx = i; q[j].idx = j }

func (q *priQueue) Push(x any) {
	it := x.(*pqItem)
	it.idx = len(*q)
	*q = append(*q, it)
}

func (q *priQueue) Pop() any {
	old := *q
	it := old[len(old)-1]
	old[len(old)-1] = nil
	*q = old[:len(old)-1]
	return it
}
