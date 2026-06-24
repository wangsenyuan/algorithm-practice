package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var q int
	fmt.Fscan(reader, &q)
	queries := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(queries)
}

const X = 1e6 + 10

var lpf [X]int
var mpf [X]int
var goodNums []int

func init() {
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			mpf[i] = i
			for j := 2 * i; j < X; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
				mpf[j] = i
			}
		}
	}

	for i := 6; i*i <= 1e12; i++ {
		x := lpf[i]
		y := mpf[i]
		if x == y {
			continue
		}
		ok := true
		for j := i; j > 1; j /= lpf[j] {
			if lpf[j] != x && lpf[j] != y {
				ok = false
				break
			}
		}
		if ok {
			goodNums = append(goodNums, i*i)
		}
	}
}

func solve(queries []int) []int {
	ans := make([]int, len(queries))

	for i, x := range queries {
		j := sort.SearchInts(goodNums, x)
		if j == len(goodNums) || goodNums[j] != x {
			j--
		}
		ans[i] = goodNums[j]
	}
	return ans
}
