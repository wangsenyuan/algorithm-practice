package main

import (
	"bufio"
	"fmt"
	"os"
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

func drive(reader *bufio.Reader) int {
	readString(reader)
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

type pair struct {
	first  int
	second int
}

func process(s string) []int {
	var arr []pair
	ss := strings.Split(s, " ")
	for _, x := range ss {
		var cnt int
		var j int
		for j < len(x) && x[j] != '-' {
			cnt = cnt*10 + int(x[j]-'0')
			j++
		}
		j++
		y := int(x[j] - 'a')
		if len(arr) == 0 || arr[len(arr)-1].second != y {
			arr = append(arr, pair{cnt, y})
		} else {
			arr[len(arr)-1].first += cnt
		}
	}
	res := make([]int, len(arr))
	for i, cur := range arr {
		res[i] = cur.first*26 + cur.second
	}
	return res
}

func solve(s string, t string) int {
	a := process(s)
	b := process(t)
	if len(b) == 1 {
		return solve1(a, b[0])
	}
	if len(b) == 2 {
		return solve2(a, b)
	}
	// main part
	m := len(b)
	first := b[0]
	last := b[m-1]
	b = b[1 : m-1]
	next := kmp(b)

	n := len(a)

	m -= 2

	var res int

	check := func(u int, v int) bool {
		return u%26 == v%26 && u/26 >= v/26
	}

	var j int
	for i := 0; i < n; i++ {
		for j > 0 && a[i] != b[j] {
			j = next[j-1]
		}
		if a[i] == b[j] {
			j++
		}
		if j == m {
			if i-m >= 0 && check(a[i-m], first) && i+1 < n && check(a[i+1], last) {
				res++
			}
			j = next[j-1]
		}
	}

	return res
}

func kmp(a []int) []int {
	n := len(a)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && a[i] != a[j] {
			j = p[j-1]
		}
		if a[i] == a[j] {
			j++
		}
		p[i] = j
	}
	return p
}

func solve1(a []int, b int) int {
	var res int
	u := b / 26
	v := b % 26
	for _, x := range a {
		if x%26 == v {
			w := x / 26
			res += max(0, w-u+1)
		}
	}
	return res
}

func solve2(a []int, b []int) int {
	var res int
	for i := 0; i+1 < len(a); i++ {
		if a[i]%26 == b[0]%26 && a[i]/26 >= b[0]/26 {
			if a[i+1]%26 == b[1]%26 && a[i+1]/26 >= b[1]/26 {
				res++
			}
		}
	}
	return res
}
