package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var x, k int
	fmt.Fscan(reader, &x, &k)
	return solve(x, k)
}

var C = 100000

func solve(x int, k int) []int {

	var res []int

	var f func(x int, k int)

	var divs []int

	for i := 1; i <= x/i; i++ {
		if x%i == 0 {
			divs = append(divs, i)
			if i*i != x {
				divs = append(divs, x/i)
			}
		}
	}

	slices.Sort(divs)
	m := len(divs)

	fs := make([][]int, m)
	for i := range m {
		for j := range i + 1 {
			if divs[i]%divs[j] == 0 {
				fs[i] = append(fs[i], divs[j])
			}
		}
	}

	f = func(x int, k int) {
		if k == 0 || x == 1 {
			res = append(res, x)
			return
		}

		i := sort.SearchInts(divs, x)
		divs := fs[i]

		if len(divs) == 2 {
			// 这个是质数
			need := min(k+1, C-len(res))
			for i := range need {
				if i < k {
					res = append(res, 1)
				} else {
					res = append(res, x)
				}
			}
			return
		}

		// 不是质数
		for _, d := range divs {
			f(d, k-1)
			if len(res) >= C {
				return
			}
		}
	}

	f(x, k)

	if len(res) > C {
		res = res[:C]
	}

	return res
}
