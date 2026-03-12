package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	var buf bytes.Buffer
	for _, cur := range res {
		buf.WriteString(strconv.Itoa(cur[0]))
		buf.WriteByte(' ')
		buf.WriteString(strconv.Itoa(cur[1]))
		buf.WriteByte('\n')
	}
	writer.Write(buf.Bytes())
}

func drive(reader *bufio.Reader) [][2]int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	customers := make([][2]int, k)
	for i := 0; i < k; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		customers[i] = [2]int{x, y}
	}
	return solve(n, m, customers)
}

func solve(n int, m int, customers [][2]int) [][2]int {
	rows := make([]rowSet, n+1)
	for i := 1; i <= n; i++ {
		rows[i] = newRowSet(m)
	}

	findSeat := func(x int, y int) (int, int) {
		bestDist := inf
		bestX, bestY := inf, inf

		update := func(r int, c int) {
			if r < 1 || r > n || c < 1 || c > m {
				return
			}
			dist := abs(r-x) + abs(c-y)
			if dist < bestDist || dist == bestDist && (r < bestX || r == bestX && c < bestY) {
				bestDist = dist
				bestX, bestY = r, c
			}
		}

		checkRow := func(r int) {
			if r < 1 || r > n {
				return
			}
			left := rows[r].findLeft(y)
			if left >= 1 {
				update(r, left)
			}
			right := rows[r].findRight(y)
			if right <= m && right != left {
				update(r, right)
			}
		}

		for d := 0; d <= bestDist && d < n; d++ {
			checkRow(x - d)
			if d > 0 {
				checkRow(x + d)
			}
		}

		return bestX, bestY
	}

	ans := make([][2]int, len(customers))
	for i, cur := range customers {
		x, y := findSeat(cur[0], cur[1])
		rows[x].erase(y)
		ans[i] = [2]int{x, y}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const inf = 1 << 60

type rowSet struct {
	left  []int
	right []int
}

func newRowSet(m int) rowSet {
	left := make([]int, m+2)
	right := make([]int, m+2)
	for i := 0; i <= m+1; i++ {
		left[i] = i
		right[i] = i
	}
	return rowSet{left: left, right: right}
}

func (rs *rowSet) findLeft(u int) int {
	p := u
	for rs.left[p] != p {
		p = rs.left[p]
	}

	for p != u {
		rs.left[u], u = p, rs.left[u]
	}

	return p
}

func (rs *rowSet) findRight(u int) int {
	p := u
	for rs.right[p] != p {
		p = rs.right[p]
	}

	for p != u {
		rs.right[u], u = p, rs.right[u]
	}

	return p
}

func (rs *rowSet) erase(pos int) {
	rs.right[pos] = rs.findRight(pos + 1)
	rs.left[pos] = rs.findLeft(pos - 1)
}
