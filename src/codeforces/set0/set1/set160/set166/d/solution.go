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
	_, _, tot, assign := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, tot)

	fmt.Fprintln(writer, len(assign))

	for _, cur := range assign {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) (shoes [][]int, customers [][]int, tot int, assign [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	shoes = make([][]int, n)
	for i := range n {
		shoes[i] = make([]int, 2)
		fmt.Fscan(reader, &shoes[i][0], &shoes[i][1])
	}

	var m int
	fmt.Fscan(reader, &m)

	customers = make([][]int, m)

	for i := range m {
		customers[i] = make([]int, 2)
		fmt.Fscan(reader, &customers[i][0], &customers[i][1])
	}

	tot, assign = solve(shoes, customers)

	return
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

type from struct {
	prev int
	cid  int
	sid  int
}

func solve(shoes [][]int, customers [][]int) (tot int, assign [][]int) {
	n := len(shoes)
	var arr []int
	for i := range n {
		arr = append(arr, shoes[i][1])
	}
	for _, cur := range customers {
		arr = append(arr, cur[1])
	}

	slices.Sort(arr)

	arr = slices.Compact(arr)

	m := len(arr)
	best := make([][]pair, m)
	for id, cur := range customers {
		d, l := cur[0], cur[1]
		i := sort.SearchInts(arr, l)
		tmp := pair{d, id}
		for j, v := range best[i] {
			if tmp.first >= v.first {
				tmp, best[i][j] = best[i][j], tmp
			}
		}
		if len(best[i]) < 2 {
			best[i] = append(best[i], tmp)
		}
	}

	price := make([]int, m)
	for id, cur := range shoes {
		i := sort.SearchInts(arr, cur[1])
		price[i] = id + 1
	}

	dp := []int{0, -inf, -inf, -inf}
	fp := make([][4]from, m)
	bestMask := 0

	for i := range m {
		ndp := []int{-inf, -inf, -inf, -inf}

		for j := 0; j < 4; j++ {
			fp[i][j] = from{-1, -1, -1}
		}

		relax := func(mask int, val int, prev int, cid int, sid int) {
			if val > ndp[mask] {
				ndp[mask] = val
				fp[i][mask] = from{prev, cid, sid}
			}
		}

		for mask, cur := range dp {
			if cur < 0 {
				continue
			}

			relax(0, cur, mask, -1, -1)

			if price[i] == 0 {
				continue
			}

			id := price[i] - 1
			c := shoes[id][0]

			for j, v := range best[i] {
				if v.first >= c {
					relax(1<<j, cur+c, mask, v.second, id)
				}
			}

			if i > 0 && arr[i-1] == arr[i]-1 {
				for j, v := range best[i-1] {
					if mask&(1<<j) == 0 && v.first >= c {
						relax(0, cur+c, mask, v.second, id)
					}
				}
			}
		}

		copy(dp, ndp)
	}

	for mask, cur := range dp {
		if cur > tot {
			tot = cur
			bestMask = mask
		}
	}

	for i := m - 1; i >= 0; i-- {
		cur := fp[i][bestMask]
		if cur.cid >= 0 {
			assign = append(assign, []int{cur.cid + 1, cur.sid + 1})
		}
		bestMask = cur.prev
	}

	return
}
