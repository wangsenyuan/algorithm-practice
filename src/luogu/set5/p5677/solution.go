package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	n, m := readInt(reader), readInt(reader)
	a := make([]int, n)
	for i := range a {
		a[i] = readInt(reader)
	}
	queries := make([][2]int, m)
	for i := range queries {
		queries[i] = [2]int{readInt(reader), readInt(reader)}
	}
	return solve(a, queries)
}

func readInt(reader *bufio.Reader) int {
	b, _ := reader.ReadByte()
	for b < '0' || b > '9' {
		b, _ = reader.ReadByte()
	}
	value := 0
	for b >= '0' && b <= '9' {
		value = value*10 + int(b-'0')
		b, _ = reader.ReadByte()
	}
	return value
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func solve(a []int, queries [][2]int) int64 {
	n := len(a)
	nums := make([]pair, n)
	for i, v := range a {
		nums[i] = pair{v, i}
	}

	slices.SortFunc(nums, func(x pair, y pair) int {
		return x.first - y.first
	})

	at := make([][]int, n)
	for i, cur := range queries {
		r := cur[1] - 1
		at[r] = append(at[r], i)
	}
	todo := make([][]int, n)

	f := make(fenwick, n+1)

	update := func(i int, k int) {
		if k < i {
			f.add(k, 1)
		} else {
			todo[k] = append(todo[k], i)
		}
	}

	var ans int

	for i, v := range a {
		j, _ := slices.BinarySearchFunc(nums, pair{v, i}, func(x pair, y pair) int {
			return x.first - y.first
		})
		dist := inf
		if j+1 < n {
			dist = nums[j+1].first - v
		}
		if j > 0 {
			dist = min(dist, v-nums[j-1].first)
		}
		if j+1 < n && nums[j+1].first-v == dist {
			update(i, nums[j+1].second)
		}
		if j > 0 && v-nums[j-1].first == dist {
			update(i, nums[j-1].second)
		}
		for _, k := range todo[i] {
			f.add(k, 1)
		}
		for _, j := range at[i] {
			l := queries[j][0] - 1
			ans += f.rangeSum(l, i) * (j + 1)
		}
	}
	return int64(ans)
}

type fenwick []int

func (f fenwick) add(i int, v int) {
	i++
	for i < len(f) {
		f[i] += v
		i += i & -i
	}
}

func (f fenwick) pre(i int) int {
	i++
	s := 0
	for i > 0 {
		s += f[i]
		i -= i & -i
	}
	return s
}

func (f fenwick) rangeSum(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}
