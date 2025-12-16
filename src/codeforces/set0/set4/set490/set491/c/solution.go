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
	_, _, _, best, res := drive(reader)
	fmt.Println(best)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) (k int, s string, t string, best int, res string) {
	first := readString(reader)
	ss := strings.Split(first, " ")
	k, _ = strconv.Atoi(ss[1])
	s = readString(reader)
	t = readString(reader)
	best, res = solve(k, s, t)
	return
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func id(x byte) int {
	if x >= 'a' && x <= 'z' {
		return int(x - 'a')
	}
	return int(x-'A') + 26
}

func solve(k int, s string, t string) (int, string) {
	a := make([][]int, k+1)
	for i := range k + 1 {
		a[i] = make([]int, k+1)
	}
	n := len(s)
	for i := range n {
		x := id(s[i])
		y := id(t[i])
		// 如果y替换成了x
		a[x+1][y+1]++
	}

	// reverse to get min value (which is max)
	for i := range k + 1 {
		for j := range k + 1 {
			a[i][j] *= -1
		}
	}

	u := make([]int, k+1)
	v := make([]int, k+1)
	p := make([]int, k+1)
	way := make([]int, k+1)
	used := make([]bool, k+1)

	minv := make([]int, k+1)

	for i := 1; i <= k; i++ {
		p[0] = i
		var j0 int

		for j := range k + 1 {
			minv[j] = inf
		}
		clear(used)
		for {
			used[j0] = true
			i0 := p[j0]
			delta := inf
			var j1 int

			for j := 1; j <= k; j++ {
				if !used[j] {
					cur := a[i0][j] - u[i0] - v[j]
					if cur < minv[j] {
						minv[j] = cur
						way[j] = j0
					}
					if minv[j] < delta {
						delta = minv[j]
						j1 = j
					}
				}
			}

			for j := range k + 1 {
				if used[j] {
					u[p[j]] += delta
					v[j] -= delta
				} else {
					minv[j] -= delta
				}
			}
			j0 = j1
			if p[j0] == 0 {
				break
			}
		}
		for {
			j1 := way[j0]
			p[j0], j0 = p[j1], j1
			if j0 == 0 {
				break
			}
		}
	}

	buf := make([]byte, k)
	for j := 1; j <= k; j++ {
		i := p[j]
		buf[i-1] = letters[j-1]
	}

	return v[0], string(buf)
}

const inf = 1 << 60
