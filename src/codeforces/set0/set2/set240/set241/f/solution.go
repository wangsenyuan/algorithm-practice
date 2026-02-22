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
	fmt.Println(res[0], res[1])
}

func readString(r *bufio.Reader) string {
	bs, _ := r.ReadBytes('\n')
	return strings.TrimSpace(string(bs))
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	nums := make([]int, len(ss))
	for i := range nums {
		nums[i], _ = strconv.Atoi(ss[i])
	}
	return nums
}

func drive(reader *bufio.Reader) []int {
	first := readNums(reader)
	m, k := first[0], first[2]
	a := make([]string, m)
	for i := range m {
		a[i] = readString(reader)
	}
	last := readString(reader)
	ss := strings.Split(last, " ")
	rs, _ := strconv.Atoi(ss[0])
	cs, _ := strconv.Atoi(ss[1])
	start := []int{rs, cs}
	path := ss[2]
	rf, _ := strconv.Atoi(ss[3])
	cf, _ := strconv.Atoi(ss[4])
	end := []int{rf, cf}

	return solve(k, a, start, path, end)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(k int, a []string, start []int, path string, end []int) []int {
	m := len(a)
	n := len(a[0])

	getCost := func(i int, j int) int {
		if a[i][j] >= 'a' && a[i][j] <= 'z' {
			return 1
		}
		return int(a[i][j] - '0')
	}

	pos := make([][]int, 26)
	for i := range m {
		for j := range n {
			if a[i][j] >= 'a' && a[i][j] <= 'z' {
				x := int(a[i][j] - 'a')
				pos[x] = []int{i, j}
			}
		}
	}

	var arr [][]int
	arr = append(arr, []int{start[0] - 1, start[1] - 1})

	for i := range path {
		x := int(path[i] - 'a')
		arr = append(arr, pos[x])
	}

	arr = append(arr, []int{end[0] - 1, end[1] - 1})

	var res [][]int

	for i := 1; i < len(arr); i++ {
		r, c := arr[i-1][0], arr[i-1][1]
		x, y := arr[i][0], arr[i][1]
		if r == x {
			d := 1
			if c > y {
				d = -1
			}
			for c != y {
				w := getCost(r, c)
				for range w {
					res = append(res, []int{r + 1, c + 1})
				}
				c += d
			}
		} else {
			// c = y
			d := 1
			if r > x {
				d = -1
			}
			for r != x {
				w := getCost(r, c)
				for range w {
					res = append(res, []int{r + 1, c + 1})
				}
				r += d
			}
		}
	}
	if len(res) <= k {
		return end
	}
	return res[k]
}

const inf = 1 << 60
