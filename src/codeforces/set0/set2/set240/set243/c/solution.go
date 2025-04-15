package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	moves := make([]string, n)
	for i := 0; i < n; i++ {
		moves[i] = readString(reader)
	}
	return solve(moves)
}

const inf = 1e10

var dd = []int{-1, 0, 1, 0, -1}

func solve(moves []string) int {
	n := len(moves)

	var xs, ys []int
	xs = append(xs, -inf, inf, 0, 1)
	ys = append(ys, -inf, inf, 0, 1)
	var x, y int
	steps := make([][]int, n)

	for i := range n {
		var v int
		readInt([]byte(moves[i]), 2, &v)
		if moves[i][0] == 'U' {
			y += v
			ys = append(ys, y, y+1)
		} else if moves[i][0] == 'D' {
			y -= v
			ys = append(ys, y, y+1)
		} else if moves[i][0] == 'R' {
			x += v
			xs = append(xs, x, x+1)
		} else {
			x -= v
			xs = append(xs, x, x+1)
		}
		steps[i] = []int{x, y}
	}
	xs = sortAndUnique(xs)
	ys = sortAndUnique(ys)

	m := len(xs)
	k := len(ys)
	color := make([][]int, m)
	for i := range m {
		color[i] = make([]int, k)
	}

	pos := make([]int, 2)
	pos[0] = sort.SearchInts(xs, 0)
	pos[1] = sort.SearchInts(ys, 0)
	tmp := make([]int, 2)

	for i := range n {
		nx, ny := steps[i][0], steps[i][1]
		tmp[0] = sort.SearchInts(xs, nx)
		tmp[1] = sort.SearchInts(ys, ny)
		if tmp[0] == pos[0] {
			u, v := pos[1], tmp[1]
			if u > v {
				u, v = v, u
			}
			for j := u; j <= v; j++ {
				color[pos[0]][j] = 1
			}
		} else {
			u, v := pos[0], tmp[0]
			if u > v {
				u, v = v, u
			}
			for j := u; j <= v; j++ {
				color[j][pos[1]] = 1
			}
		}
		copy(pos, tmp)
	}
	// 从(0, 0), (m-1, k - 1)开始访问
	que := make([]int, m*k)
	var head, tail int
	que[head] = 0
	head++
	que[head] = m*k - 1
	head++
	color[0][0] = 2
	color[m-1][k-1] = 2
	for tail < head {
		r, c := que[tail]/k, que[tail]%k
		tail++
		for i := 0; i < 4; i++ {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < m && nc >= 0 && nc < k && color[nr][nc] == 0 {
				que[head] = nr*k + nc
				head++
				color[nr][nc] = 2
			}
		}
	}

	var ans int
	for i := range m {
		for j := range k {
			if color[i][j] <= 1 {
				dx := xs[i+1] - xs[i]
				dy := ys[j+1] - ys[j]
				ans += dx * dy
			}
		}
	}
	return ans
}

func sortAndUnique(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	sort.Ints(res)
	var n int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] > res[i-1] {
			res[n] = res[i-1]
			n++
		}
	}
	return res[:n]
}
