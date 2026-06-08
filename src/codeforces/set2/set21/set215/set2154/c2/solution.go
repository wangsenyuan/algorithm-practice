package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

const N = 200010

var lpf [N]int
var primes []int

func init() {
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= N {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) int {

	best := make([][]int, len(primes))

	addBest := func(f int, v int) {
		for i, w := range best[f] {
			if v <= w {
				v, best[f][i] = best[f][i], v
			}
		}
		if len(best[f]) < 2 {
			best[f] = append(best[f], v)
		}
	}

	freq := make([]int, len(primes))

	for _, u := range a {
		// v := b[i]
		if u > 1 {
			for u > 1 {
				x := lpf[u]
				for u%x == 0 {
					u /= x
				}
				j := sort.SearchInts(primes, x)
				freq[j]++
				if freq[j] > 1 {
					return 0
				}
			}
		}
	}

	min_cost_index := -1

	for i, u := range a {
		v := b[i]

		if min_cost_index < 0 || v < b[min_cost_index] {
			min_cost_index = i
		}

		if u == 1 {
			addBest(0, v)
		} else {
			u++
			for u > 1 {
				x := lpf[u]
				for u%x == 0 {
					u /= x
				}
				j := sort.SearchInts(primes, x)
				addBest(j, v)
			}
		}
	}

	ans := 1 << 60

	for i, cur := range best {

		if len(cur) > 1 {
			ans = min(ans, (cur[0] + cur[1]))
		}
		if len(cur) > 0 && freq[i] > 0 {
			ans = min(ans, cur[0])
		}
	}

	for u := a[min_cost_index]; u > 1; {
		x := lpf[u]
		for u%x == 0 {
			u /= x
		}
		freq[sort.SearchInts(primes, x)]--
	}

	for i, x := range primes {
		if freq[i] > 0 {
			// 存在一个整除x的数, num + w = m * x
			num := a[min_cost_index]
			diff := (num+x-1)/x*x - num
			ans = min(ans, diff*b[min_cost_index])
		}
	}

	return ans
}
