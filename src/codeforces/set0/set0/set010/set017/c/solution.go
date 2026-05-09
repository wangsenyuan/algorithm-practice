package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	s := readString(reader)
	if len(s) > n {
		s = s[:n]
	}
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const mod = 51123987

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(s string) int {
	n := len(s)
	arr := make([]int, n)
	for i := range n {
		arr[i] = int(s[i] - 'a')
	}

	arr = slices.Compact(arr)
	m := len(arr)
	next := make([][]int, m+1)
	pos := []int{m, m, m}
	for i := m - 1; i >= 0; i-- {
		pos[arr[i]] = i
		next[i] = slices.Clone(pos)
	}
	next[m] = []int{m, m, m}

	h := (n + 2) / 3
	strideB := m * 3
	strideA := (h + 1) * strideB
	size := (h + 1) * strideA
	id := func(a int, b int, i int, last int) int {
		return a*strideA + b*strideB + i*3 + last
	}

	cur := make([]int, size)
	for x := range 3 {
		if next[0][x] == m {
			continue
		}
		switch x {
		case 0:
			cur[id(1, 0, next[0][x], x)] = 1
		case 1:
			cur[id(0, 1, next[0][x], x)] = 1
		default:
			cur[id(0, 0, next[0][x], x)] = 1
		}
	}
	nxt := make([]int, size)

	for l := 1; l < n; l++ {
		for a := 0; a <= h; a++ {
			for b := 0; b <= h; b++ {
				c := l - a - b
				if c < 0 || c > h {
					continue
				}
				for i := range m {
					for last := range 3 {
						v := cur[id(a, b, i, last)]
						if v == 0 {
							continue
						}
						for x := range 3 {
							na, nb, nc := a, b, c
							switch x {
							case 0:
								na++
							case 1:
								nb++
							default:
								nc++
							}
							if max(na, nb, nc) > h {
								continue
							}
							j := i
							if x != last {
								j = next[i+1][x]
								if j == m {
									continue
								}
							}
							k := id(na, nb, j, x)
							nxt[k] = add(nxt[k], v)
						}
					}
				}
			}
		}
		copy(cur, nxt)
		clear(nxt)
	}

	var res int
	for a := 0; a <= h; a++ {
		for b := 0; b <= h; b++ {
			c := n - a - b
			if c < 0 || c > h || max(a, b, c)-min(a, b, c) > 1 {
				continue
			}
			for i := range m {
				for last := range 3 {
					res = add(res, cur[id(a, b, i, last)])
				}
			}
		}
	}
	return res
}
