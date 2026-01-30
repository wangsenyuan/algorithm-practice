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

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func drive(reader *bufio.Reader) int {
	first := readString(reader)
	ss := strings.Split(first, " ")
	n := parseInt(ss[0])

	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(grid []string) int {
	n := len(grid)
	m := len(grid[0])

	check := func(flag [][27]int) int {
		// 可以用O(n)找出所有的回文
		w := findPalindrome(flag)
		var ans int
		for i := range n {
			if flag[i][26] < 0 {
				continue
			}
			ans += (w[2*i+1] + 1) / 2
			if i > 0 {
				ans += w[2*i] / 2
			}
		}
		return ans
	}

	var res int

	for c1 := range m {
		flag := make([][27]int, n)
		odd := make([]int, n)
		for c2 := c1; c2 < m; c2++ {
			var neg int
			for r := range n {
				x := int(grid[r][c2] - 'a')
				flag[r][x]++
				if flag[r][x]&1 == 1 {
					odd[r]++
				} else {
					odd[r]--
				}

				if odd[r] > 1 {
					neg--
					flag[r][26] = neg
				} else {
					flag[r][26] = 0
				}
			}
			res += check(flag)
		}
	}

	return res
}

func findPalindrome(seq [][27]int) []int {
	n := len(seq)
	var l []int
	var palLen int
	for i := 0; i < n; {
		if i > palLen && seq[i-palLen-1] == seq[i] {
			palLen += 2
			i++
			continue
		}

		l = append(l, palLen)

		s := len(l) - 2
		e := s - palLen
		found := false
		for j := s; j > e; j-- {

			d := j - e - 1
			if l[j] == d {
				found = true
				palLen = d
				break
			}

			l = append(l, min(d, l[j]))
		}
		if !found {
			palLen = 1
			i++
		}
	}

	l = append(l, palLen)
	s := len(l) - 2
	e := s - (2*n + 1 - len(l))
	for i := s; i > e; i-- {
		d := i - e - 1
		l = append(l, min(d, l[i]))
	}
	return l
}

func solve1(grid []string) int {
	n := len(grid)
	m := len(grid[0])

	sz := n*2 + 3
	halfLen := make([]int, sz-2)
	halfLen[1] = 1

	var res int

	for r := range m {
		t := make([][26]int, sz)
		odd := make([]int, sz)
		odd[0] = 2
		odd[sz-1] = 2

		for l := r; l >= 0; l-- {
			for i, row := range grid {
				x := int(row[l] - 'a')
				t[i*2+2][x]++
				odd[i*2+2] += t[i*2+2][x]%2*2 - 1
			}

			boxM, boxR := 0, 0
			for i := 2; i < sz-2; i++ {
				hl := 0
				if i < boxR {
					hl = min(halfLen[boxM*2-i], boxR-i)
				}
				for odd[i-hl] <= 1 && odd[i+hl] <= 1 && t[i-hl] == t[i+hl] {
					hl++
					boxM, boxR = i, i+hl
				}
				halfLen[i] = hl
				res += hl / 2
			}
		}
	}

	return res
}
