package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/bits"
	"os"
	"slices"
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

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	arr := strings.Split(s, " ")
	res := make([]int, len(arr))
	for i := range len(arr) {
		res[i], _ = strconv.Atoi(arr[i])
	}
	return res
}

func drive(reader *bufio.Reader) string {
	first := readNums(reader)
	n, k := first[0], first[2]
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(k, a)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(k int, a []string) string {
	n := len(a)
	m := len(a[0])

	var start []int
	var dest []int

	var letters []int
	for i, row := range a {
		for j, c := range row {
			if c >= 'a' && c <= 'z' {
				letters = append(letters, int(c-'a'))
			} else if c == 'S' {
				start = append(start, i, j)
			} else {
				dest = append(dest, i, j)
			}
		}
	}

	if abs(start[0]-dest[0])+abs(start[1]-dest[1]) == 1 {
		return ""
	}

	letters = sortAndUnique(letters)

	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, m)
	}
	que := make([]int, n*m)

	var ans string

	bfs := func(mask int) string {
		for i := range n {
			for j := range m {
				dist[i][j] = -1
			}
		}
		var head, tail int
		que[head] = dest[0]*m + dest[1]
		head++
		dist[dest[0]][dest[1]] = 0
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			if len(ans) > 0 && dist[r][c] > len(ans) || r == start[0] && c == start[1] {
				// 剪枝
				break
			}
			for i := range 4 {
				x, y := r+dd[i], c+dd[i+1]
				if x >= 0 && x < n && y >= 0 && y < m && dist[x][y] < 0 && (a[x][y] == 'S' || (mask&(1<<(a[x][y]-'a'))) > 0) {
					dist[x][y] = dist[r][c] + 1
					que[head] = x*m + y
					head++
				}
			}
		}
		if dist[start[0]][start[1]] < 0 {
			return ""
		}

		buf := make([]byte, dist[start[0]][start[1]]+1)
		head = 0
		tail = 0
		que[head] = start[0]*m + start[1]
		head++
		for tail < head {
			cur := head
			w := dist[que[tail]/m][que[tail]%m]
			buf[w] = a[que[tail]/m][que[tail]%m]
			if w == 0 {
				break
			}
			var next [][2]int
			for tail < cur {
				r, c := que[tail]/m, que[tail]%m
				tail++
				for i := range 4 {
					nr, nc := r+dd[i], c+dd[i+1]
					if nr >= 0 && nr < n && nc >= 0 && nc < m && dist[nr][nc] == w-1 {
						next = append(next, [2]int{nr, nc})
					}
				}
			}
			slices.SortFunc(next, func(f [2]int, s [2]int) int {
				if a[f[0]][f[1]] < a[s[0]][s[1]] {
					return -1
				}
				if a[f[0]][f[1]] > a[s[0]][s[1]] {
					return 1
				}
				return cmp.Or(f[0]-s[0], f[1]-s[1])
			})

			next = slices.Compact(next)

			ch := a[next[0][0]][next[0][1]]

			for _, cur := range next {
				r, c := cur[0], cur[1]
				if a[r][c] != ch {
					break
				}
				que[head] = r*m + c
				head++
			}
		}
		slices.Reverse(buf)
		return string(buf[1 : len(buf)-1])
	}

	update := func(tmp string) {
		if len(tmp) == 0 {
			return
		}
		if len(ans) == 0 || len(tmp) < len(ans) || len(tmp) == len(ans) && tmp < ans {
			ans = tmp
		}
	}

	k = min(k, len(letters))
	state := 1<<k - 1
	for {
		var flag int
		for i := range len(letters) {
			if (state>>i)&1 == 1 {
				flag |= 1 << letters[i]
			}
		}
		tmp := bfs(flag)
		update(tmp)
		state = nextPermuatation(state)
		if bits.Len(uint(state)) > len(letters) {
			break
		}
	}

	if len(ans) == 0 {
		return "-1"
	}

	return ans
}

func sortAndUnique(a []int) []int {
	slices.Sort(a)
	return slices.Compact(a)
}

func nextPermuatation(state int) int {
	if state == 0 {
		return 0
	}
	lowest := state & -state
	next := state + lowest
	ones := state ^ next
	ones = (ones / lowest) >> 2
	return next | ones
}

func abs(num int) int {
	return max(num, -num)
}
