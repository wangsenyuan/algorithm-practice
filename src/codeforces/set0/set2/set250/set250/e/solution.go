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
	res := process(reader)
	if res < 0 {
		fmt.Println("Never")
	} else {
		fmt.Println(res)
	}
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	return strings.TrimSpace(string(s))
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)
	}
	return solve(n, m, grid)
}

const inf = 1 << 30

func solve(n int, m int, grid []string) int {

	buf := make([]byte, m+2)
	buf[0] = '#'
	buf[m+1] = '#'

	t0 := NewSegTree(m+2, inf, min)
	t1 := NewSegTree(m+2, -inf, max)

	marked := make([]int, m+2)

	move := func(r int, c int, d int) (int, int, int) {
		// 在(r, c)处进入第r行，方向是d(0 for right, 1 for left)
		// 如果能到达下层，返回位置和方向
		defer func() {
			clear(marked)
			t0.Reset()
			t1.Reset()
		}()

		copy(buf[1:], grid[r])
		var pos []int

		for i := 0; i <= m+1; i++ {
			if buf[i] != '.' {
				t0.Update(i, i)
				t1.Update(i, i)
			}
			// 可以降落的位置
			if i >= 1 && i <= m && grid[r+1][i-1] == '.' {
				pos = append(pos, i)
			}
		}

		if len(pos) == 0 {
			return -1, 0, 0
		}

		c++
		pc := sort.SearchInts(pos, c)
		if pc < len(pos) && pos[pc] == c {
			// drop dwon
			return 1, c - 1, d
		}

		var steps int

		for {
			if d == 0 {
				nc := t0.Query(c+1, m+2)
				pc = sort.SearchInts(pos, c)
				if pc < len(pos) && pos[pc] < nc {
					steps += pos[pc] - c
					return steps + 1, pos[pc] - 1, d
				}
				steps += nc - c
				if buf[nc] == '+' {
					t0.Update(nc, inf)
					t1.Update(nc, -inf)
					buf[nc] = '.'
				} else {
					// buf[nc] = '#'
					if marked[nc] == m {
						return -1, 0, 0
					}
					// 又遇到了这面墙
					marked[nc]++
				}
				c = nc - 1
			} else {
				nc := t1.Query(0, c)
				pc = sort.SearchInts(pos, c)
				if pc == len(pos) || pos[pc] > c {
					pc--
				}
				// pos[pc] <= c
				if pc >= 0 && nc < pos[pc] {
					steps += c - pos[pc]
					return steps + 1, pos[pc] - 1, d
				}
				steps += c - nc
				if buf[nc] == '+' {
					t0.Update(nc, inf)
					t1.Update(nc, -inf)
					buf[nc] = '.'
				} else {
					if marked[nc] == m {
						return -1, 0, 0
					}
					marked[nc]++
				}
				c = nc + 1
			}

			d ^= 1
		}
	}
	var steps int
	var c, d int
	for r := 0; r+1 < n; r++ {
		tmp, nc, nd := move(r, c, d)
		if tmp < 0 {
			return -1
		}
		steps += tmp
		c = nc
		d = nd
	}

	return steps
}

func abs(num int) int {
	return max(num, -num)
}

type SegTree struct {
	f   func(int, int) int
	arr []int
	iv  int
	n   int
}

func NewSegTree(n int, iv int, f func(int, int) int) *SegTree {
	arr := make([]int, n*2)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{f, arr, iv, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] = v

	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Query(l int, r int) int {
	res := t.iv
	l += t.n
	r += t.n
	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func (t *SegTree) Reset() {
	for i := range len(t.arr) {
		t.arr[i] = t.iv
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func bruteForce(n int, m int, grid []string) int {
	var steps int
	var c, d int
	marked := make([]int, m+2)
	buf := make([][]byte, n)
	for i := range n {
		buf[i] = []byte(grid[i])
	}
	for r := 0; r+1 < n; r++ {
		for {
			if buf[r+1][c] == '.' {
				steps++
				break
			}
			if d == 0 {
				nc := c + 1
				if nc == m || buf[r][nc] == '#' {
					if marked[nc+1] == m {
						return -1
					}
					marked[nc+1]++
					c = nc - 1
					d ^= 1
				} else if buf[r][nc] == '+' {
					buf[r][nc] = '.'
					c = nc - 1
					d ^= 1
				} else {
					c = nc
				}
			} else {
				nc := c - 1
				if nc < 0 || buf[r][nc] == '#' {
					if marked[nc+1] == m {
						return -1
					}
					marked[nc+1]++
					c = nc + 1
					d ^= 1
				} else if buf[r][nc] == '+' {
					buf[r][nc] = '.'
					c = nc + 1
					d ^= 1
				} else {
					c = nc
				}
			}
			steps++
		}
		clear(marked)
	}
	return steps
}
