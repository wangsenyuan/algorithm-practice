package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	cnt := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &cnt[i])
	}
	price := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &price[i])
	}
	var r int
	fmt.Fscan(reader, &r)
	return solve(s, cnt, price, r)
}

func solve(hamburger string, cnt []int, price []int, s int) int {
	freq := make([]int, 3)
	n := len(hamburger)
	for i := range n {
		switch hamburger[i] {
		case 'B':
			freq[0]++
		case 'S':
			freq[1]++
		default:
			freq[2]++
		}
	}

	check := func(x int) bool {
		a := max(x*freq[0]-cnt[0], 0)
		b := max(x*freq[1]-cnt[1], 0)
		c := max(x*freq[2]-cnt[2], 0)
		return a*price[0]+b*price[1]+c*price[2] <= s
	}

	w := slices.Min(price)
	r := s/w + 1
	for i := range 3 {
		if freq[i] > 0 {
			r += cnt[i] / freq[i]
		}
	}

	var l int
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}
