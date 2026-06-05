package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cnt, field := drive(reader)

	var buf bytes.Buffer
	fmt.Fprintln(&buf, cnt)
	for _, row := range field {
		fmt.Fprintln(&buf, row)
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) (int64, []string) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	field := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &field[i])
	}
	return solve(field)
}

type cell struct {
	r int
	c int
}

func solve(field []string) (int64, []string) {
	n, m := len(field), len(field[0])
	var squareId int
	square := make([][]int, n)
	col := make([]int, m)
	for i := range n {
		square[i] = make([]int, m)
		for j := range m {
			square[i][j] = -1
		}
	}

	for i := range n {
		var row int
		for j := range m {
			if field[i][j] != '.' {
				row++
				col[j]++
			}
			if row%2 == 1 && col[j]%2 == 1 {
				for dr := 0; dr <= 1; dr++ {
					for dc := 0; dc <= 1; dc++ {
						square[i+dr][j+dc] = squareId
					}
				}
				squareId++
			}
		}
	}

	pos := make(map[byte][]cell)
	for i := range n {
		for j := range m {
			if field[i][j] != '.' {
				pos[field[i][j]] = append(pos[field[i][j]], cell{i, j})
			}
		}
	}

	var edges [][2]int
	for _, ps := range pos {
		a := square[ps[0].r][ps[0].c]
		b := square[ps[1].r][ps[1].c]
		edges = append(edges, [2]int{a, b})
	}

	pairId := make([][]int, 7)
	for i := range 7 {
		pairId[i] = make([]int, 7)
	}
	var pid int
	for i := range 7 {
		for j := i; j < 7; j++ {
			pairId[i][j] = pid
			pairId[j][i] = pid
			pid++
		}
	}

	color := make([]int, 14)
	for i := range color {
		color[i] = -1
	}
	cntColor := make([]int, 7)
	// usedPair := make([]bool, 28)
	var first []int
	var ways int64

	var dfs func(int, int, int)
	dfs = func(v int, mx int, used int) {
		if v == 14 {
			if mx != 6 {
				return
			}
			for i := range 7 {
				if cntColor[i] != 2 {
					return
				}
			}
			ways++
			if first == nil {
				first = append([]int(nil), color...)
			}
			return
		}

		lim := min(6, mx+1)
		for c := 0; c <= lim; c++ {
			if cntColor[c] == 2 {
				continue
			}

			newUsed := used
			ok := true
			for _, e := range edges {
				if e[0] != v && e[1] != v {
					continue
				}

				otherColor := c
				if e[0] != e[1] {
					other := e[0] + e[1] - v
					if color[other] < 0 {
						continue
					}
					otherColor = color[other]
				}

				p := pairId[c][otherColor]
				if (newUsed>>p)&1 == 1 {
					ok = false
					break
				}
				newUsed |= 1 << p
			}

			if ok {
				color[v] = c
				cntColor[c]++
				dfs(v+1, max(mx, c), newUsed)
				cntColor[c]--
				color[v] = -1
			}

		}
	}

	dfs(0, -1, 0)

	ans := make([][]byte, n)
	for i := range n {
		ans[i] = []byte(field[i])
		for j := range m {
			if square[i][j] >= 0 {
				ans[i][j] = byte('0' + first[square[i][j]])
			}
		}
	}
	res := make([]string, n)
	for i := range n {
		res[i] = string(ans[i])
	}

	return ways * 5040, res
}
