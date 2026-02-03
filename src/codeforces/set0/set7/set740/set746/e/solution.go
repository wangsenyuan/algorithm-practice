package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, cnt, res := drive(reader)
	fmt.Println(cnt)
	if cnt < 0 {
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (m int, a []int, cnt int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &m)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	cnt, res = solve(m, a)
	return
}

func solve(m int, a []int) (int, []int) {
	n := len(a)
	if n%2 == 1 {
		return -1, nil
	}

	cnt := make([]int, 2)
	for _, x := range a {
		cnt[x%2]++
	}

	type pair struct {
		first  int
		second int
	}

	play := func(d int, diff int) []int {
		arr := slices.Clone(a)

		marked := make(map[int]bool)

		for _, x := range arr {
			if x&1 == d^1 && x <= m {
				// 这个保留下来，不要被交换掉，
				marked[x] = true
			}
		}

		var ps []pair
		for i, x := range arr {
			if x&1 == d {
				ps = append(ps, pair{x, i})
			}
		}

		slices.SortFunc(ps, func(a, b pair) int {
			return cmp.Or(a.first-b.first, a.second-b.second)
		})

		var pos [][]int
		for i := 0; i < len(ps); {
			j := i
			var cur []int
			for i < len(ps) && ps[i].first == ps[j].first {
				cur = append(cur, ps[i].second)
				i++
			}
			pos = append(pos, cur)
		}
		// 优先交换多的部分
		slices.SortFunc(pos, func(a []int, b []int) int {
			return len(b) - len(a)
		})

		rem := 1
		for i, p := range pos {
			// 把最后一个留下来
			for diff > 0 && len(p) > 1 {
				for rem <= m && (marked[rem] || rem&1 == d) {
					rem++
				}
				if rem > m {
					// no solution
					return nil
				}
				marked[rem] = true
				diff -= 2
				arr[p[0]] = rem
				p = p[1:]
			}
			pos[i] = p
		}
		// 优先目标diff = 0
		for i := 0; i < len(pos) && diff > 0; i++ {
			p := pos[i]
			// len(p) = 1
			for rem <= m && (marked[rem] || rem&1 == d) {
				rem++
			}
			if rem > m {
				return nil
			}
			diff -= 2
			marked[rem] = true
			arr[p[0]] = rem
		}

		if diff > 0 {
			return nil
		}
		// diff = 0
		// 处理重复的数字
		for _, x := range arr {
			if x <= m {
				// 这些已经在arr中的，且 <= m 的数，不能被交换进来
				marked[x] = true
			}
		}

		vis := make(map[int]bool)
		val := []int{2, 1}
		for i := range arr {
			x := arr[i]
			if vis[x] {
				for val[x&1] <= m && marked[val[x&1]] {
					val[x&1] += 2
				}
				if val[x&1] > m {
					return nil
				}
				marked[val[x&1]] = true
				arr[i] = val[x&1]
				val[x&1] += 2
			}
			vis[arr[i]] = true
		}

		return arr
	}

	var res []int
	if cnt[0] >= cnt[1] {
		res = play(0, cnt[0]-cnt[1])
	} else {
		res = play(1, cnt[1]-cnt[0])
	}

	if len(res) == 0 {
		return -1, nil
	}

	var diff int
	for i := range n {
		if res[i] != a[i] {
			diff++
		}
	}

	return diff, res
}
