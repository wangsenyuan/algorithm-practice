package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var k int
	fmt.Fscan(reader, &k)
	swaps := make([][2]int, k)
	for i := range swaps {
		fmt.Fscan(reader, &swaps[i][0], &swaps[i][1])
	}
	return solve(swaps)
}

func solve(swaps [][2]int) int64 {
	seen := make([]int, 0, len(swaps)*2)
	for _, swap := range swaps {
		seen = append(seen, swap[0], swap[1])
	}
	slices.Sort(seen)
	seen = slices.Compact(seen)

	pos := make(map[int]int, len(seen))
	arr := make([]int, len(seen))
	for i, v := range seen {
		pos[v] = i
		arr[i] = i
	}
	for _, swap := range swaps {
		i, j := pos[swap[0]], pos[swap[1]]
		arr[i], arr[j] = arr[j], arr[i]
	}

	tr := make(fenwick, len(arr)+2)
	var ans, placed int64
	for _, rank := range arr {
		ans += placed - tr.pre(rank)
		placed++
		tr.update(rank, 1)
	}
	for i, rank := range arr {
		if rank == i {
			continue
		}
		lo, hi := i, rank
		if rank < i {
			lo, hi = rank, i
		}
		ans += int64(seen[hi]-seen[lo]-1) - int64(hi-lo-1)
	}
	return ans
}

type fenwick []int64

func (f fenwick) update(i int, v int64) {
	i++
	for i < len(f) {
		f[i] += v
		i += i & -i
	}
}

func (f fenwick) pre(i int) int64 {
	i++
	var sum int64
	for i > 0 {
		sum += f[i]
		i -= i & -i
	}
	return sum
}
