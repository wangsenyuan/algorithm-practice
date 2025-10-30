package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	tc := readNum(reader)
	for tc > 0 {
		tc--
		res := drive(reader)
		if res {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
		if tc > 0 {
			readString(reader)
		}
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) bool {
	a := make([]string, 8)
	for i := range 8 {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) bool {
	n := len(a)
	var knights [][]int
	for i := range n {
		for j := range n {
			if a[i][j] == 'K' {
				knights = append(knights, []int{i, j})
			}
		}
	}

	var dd = [][]int{
		{-2, -2}, {-2, 2}, {2, -2}, {2, 2},
	}

	bfs := func(x int, y int) [][]int {
		dist := make([][]int, n)
		for i := range n {
			dist[i] = make([]int, n)
			for j := range n {
				dist[i][j] = -1
			}
		}
		dist[x][y] = 0
		var que [][]int
		que = append(que, []int{x, y})

		for len(que) > 0 {
			cur := que[0]
			que = que[1:]
			r, c := cur[0], cur[1]

			for _, d := range dd {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n && dist[nr][nc] == -1 {
					dist[nr][nc] = dist[r][c] + 1
					que = append(que, []int{nr, nc})
				}
			}
		}
		return dist
	}

	d1 := bfs(knights[0][0], knights[0][1])

	res := d1[knights[1][0]][knights[1][1]]
	return res > 0 && res&1 == 0
}
