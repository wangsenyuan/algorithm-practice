package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInts(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i, x := range ss {
		res[i], _ = strconv.Atoi(x)
	}
	return res
}

func drive(reader *bufio.Reader) int {
	first := readInts(reader)
	k, n := first[0], first[1]
	layers := make([][]string, k)
	for i := range k {
		readString(reader)
		layers[i] = make([]string, n)
		for j := range n {
			layers[i][j] = readString(reader)
		}
	}
	readString(reader)
	second := readInts(reader)
	return solve(layers, second[0], second[1])
}

func solve(layers [][]string, x, y int) int {
	k := len(layers)
	n := len(layers[0])
	m := len(layers[0][0])
	marked := make([][][]int, k)
	for i := range k {
		marked[i] = make([][]int, n)
		for j := range n {
			marked[i][j] = make([]int, m)
		}
	}

	var que [][]int

	que = append(que, []int{0, x - 1, y - 1})
	marked[0][x-1][y-1] = 1

	var dd = []int{-1, 0, 1, 0, -1}

	for len(que) > 0 {
		d, r, c := que[0][0], que[0][1], que[0][2]
		que = que[1:]
		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m &&
				layers[d][nr][nc] == '.' && marked[d][nr][nc] == 0 {
				marked[d][nr][nc] = 1
				que = append(que, []int{d, nr, nc})
			}
		}

		if d+1 < k && marked[d+1][r][c] == 0 && layers[d+1][r][c] == '.' {
			marked[d+1][r][c] = 1
			que = append(que, []int{d + 1, r, c})
		}
		if d > 0 && marked[d-1][r][c] == 0 && layers[d-1][r][c] == '.' {
			marked[d-1][r][c] = 1
			que = append(que, []int{d - 1, r, c})
		}
	}

	var sum int
	for d := range k {
		for i := range n {
			for j := range m {
				sum += marked[d][i][j]
			}
		}
	}

	return sum
}
