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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	sectors := make([][]int, n)
	for i := range n {
		var k int
		fmt.Fscan(reader, &k)
		sectors[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &sectors[i][j])
		}
	}
	return solve(n, sectors)
}

func solve(n int, sectors [][]int) int {

	for i := range n {
		sort.Ints(sectors[i])
	}

	get := func(x int) int {
		l := (x + 1) % n
		r := (x + n - 1) % n
		var res int
		for i, j, k := 0, 0, 0; i+1 < len(sectors[x]); i++ {
			// 需要知道l/r在区间i,i+1之间着力点的数量

			for j < len(sectors[l]) && sectors[l][j] < sectors[x][i] {
				j++
			}
			var c1 int
			if j < len(sectors[l]) && sectors[l][j] < sectors[x][i+1] {
				c1++
				for j+1 < len(sectors[l]) && sectors[l][j+1] < sectors[x][i+1] {
					j++
					c1++
				}
			}

			for k < len(sectors[r]) && sectors[r][k] < sectors[x][i] {
				k++
			}
			var c2 int
			if k < len(sectors[r]) && sectors[r][k] < sectors[x][i+1] {
				c2++
				for k+1 < len(sectors[r]) && sectors[r][k+1] < sectors[x][i+1] {
					k++
					c2++
				}
			}
			if c1 != c2 {
				res++
			}
		}
		return res
	}

	var res int

	for i := range n {
		res += get(i)
	}

	return res
}
