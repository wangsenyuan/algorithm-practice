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
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const inf = 1 << 30

type pair struct {
	first  int
	second int
}

type Num struct {
	val          int
	factors      map[int]int
	factor_count int
}

func solve(a []int) int {
	// a[i]很大
	n := len(a)
	sort.Ints(a)
	slices.Reverse(a)

	fs := make([]Num, n)
	var prime_cnt int
	for i := range n {
		fs[i] = factor(a[i])
		if fs[i].factor_count == 1 {
			prime_cnt++
		}
	}

	var tr []int
	tr = append(tr, 0)

	var dfs func(cur int, root_children_cnt int, root_children_sum int) int

	dfs = func(cur int, root_children_cnt int, root_children_sum int) int {
		if cur == n {
			res := root_children_sum
			if root_children_cnt > 1 {
				// root需要一个额外的节点
				res++
			}
			res += n
			res -= prime_cnt
			return res
		}
		best := inf

		for i := 0; i < len(tr); i++ {
			if tr[i]%a[cur] == 0 {
				tr[i] /= a[cur]
				tr = append(tr, a[cur])
				new_root_children_cnt := root_children_cnt
				new_root_children_sum := root_children_sum
				if i == 0 {
					new_root_children_cnt++
					new_root_children_sum += fs[cur].factor_count
				}
				best = min(best, dfs(cur+1, new_root_children_cnt, new_root_children_sum))
				tr = tr[:len(tr)-1]
				tr[i] *= a[cur]
			}
		}

		return best
	}

	return dfs(0, 0, 0)
}

func factor(num int) Num {
	res := Num{val: num}
	freq := make(map[int]int)

	for x := 2; x <= num/x; x++ {
		for num%x == 0 {
			freq[x]++
			num /= x
		}
	}
	if num > 1 {
		freq[num]++
	}

	res.factors = freq
	res.factor_count = sum_value(freq)
	return res
}

func sum_value(f map[int]int) int {
	res := 0
	for _, v := range f {
		res += v
	}
	return res
}
