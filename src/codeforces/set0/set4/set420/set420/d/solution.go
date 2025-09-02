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
	res := drive(reader)

	if len(res) == 0 {
		fmt.Println("-1")
		return
	}

	var result []string
	for _, v := range res {
		result = append(result, fmt.Sprintf("%d", v))
	}
	fmt.Println(strings.Join(result, " "))
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	ops := make([][]int, m)
	for i := range m {
		ops[i] = make([]int, 2)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1])
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) []int {
	var sz int
	marked := make([]bool, n+1)
	ans := make([]int, n)

	// Initialize answer array with -1
	for i := range n {
		ans[i] = -1
	}

	bit := make(BIT, n+10)

	find := func(y int) int {
		return sort.Search(y, func(j int) bool {
			return j-bit.pre(j) >= y-sz
		})
	}

	assign := func(y int, x int) bool {
		i := find(y - 1)

		if i >= y || ans[i] != -1 {
			return false
		}
		ans[i] = x
		marked[x] = true
		bit.Update(i, 1)
		return true
	}

	for _, op := range ops {
		x, y := op[0], op[1]
		if y <= sz {
			if !marked[x] {
				return nil
			}
		} else {
			if marked[x] {
				return nil
			}
			if !assign(y, x) {
				return nil
			}
			sz++
		}
	}

	// Find unused numbers and positions
	var nums, pos []int

	for i := 1; i <= n; i++ {
		if !marked[i] {
			nums = append(nums, i)
		}
	}

	for i := 0; i < n; i++ {
		if ans[i] == -1 {
			pos = append(pos, i)
		}
	}

	if len(pos) != len(nums) {
		return nil
	}

	// Fill remaining positions with unused numbers
	for i := range len(pos) {
		ans[pos[i]] = nums[i]
	}

	return ans
}

type BIT []int

func (bit BIT) Update(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) pre(p int) int {
	var res int
	p++
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	return bit.pre(r) - bit.pre(l-1)
}
