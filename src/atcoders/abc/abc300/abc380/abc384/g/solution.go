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

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func drive(reader *bufio.Reader) []int64 {
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

	var k int
	fmt.Fscan(reader, &k)
	queries := make([][2]int, k)
	for i := range k {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}

	return solve(a, b, queries)
}

type query struct {
	x  int
	y  int
	id int
}

func solve(A []int, B []int, queries [][2]int) []int64 {
	arr := make([]query, len(queries))

	for i := range queries {
		arr[i] = query{
			x:  queries[i][0],
			y:  queries[i][1],
			id: i,
		}
	}

	blockSize := 1000

	slices.SortFunc(arr, func(a query, b query) int {
		if a.x/blockSize != b.x/blockSize {
			return a.x/blockSize - b.x/blockSize
		}
		if a.x/blockSize%2 == 0 {
			return a.y - b.y
		}
		return b.y - a.y
	})

	nums := slices.Clone(A)
	nums = append(nums, B...)
	slices.Sort(nums)
	nums = slices.Compact(nums)

	a := make([]int, len(A))
	b := make([]int, len(B))

	for i := range a {
		a[i] = sort.SearchInts(nums, A[i])
		b[i] = sort.SearchInts(nums, B[i])
	}

	sz := len(nums)
	fa0 := make(fenwick, sz+3)
	fa1 := make(fenwick, sz+3)
	fb0 := make(fenwick, sz+3)
	fb1 := make(fenwick, sz+3)

	ans := make([]int64, len(queries))

	var x, y int
	var now, asum, bsum int

	for _, q := range arr {
		for x < q.x {
			c := fb0.queryRange(0, a[x])
			s := fb1.queryRange(0, a[x])

			now += y*A[x] + bsum - 2*((y-c)*A[x]+s)
			asum += A[x]
			fa0.update(a[x], 1)
			fa1.update(a[x], A[x])
			x++
		}

		for x > q.x {
			x--
			c := fb0.queryRange(0, a[x])
			s := fb1.queryRange(0, a[x])

			now -= y*A[x] + bsum - 2*((y-c)*A[x]+s)
			asum -= A[x]
			fa0.update(a[x], -1)
			fa1.update(a[x], -A[x])
		}

		for y < q.y {
			c := fa0.queryRange(0, b[y])
			s := fa1.queryRange(0, b[y])

			now += x*B[y] + asum - 2*((x-c)*B[y]+s)
			bsum += B[y]
			fb0.update(b[y], 1)
			fb1.update(b[y], B[y])
			y++
		}

		for y > q.y {
			y--
			c := fa0.queryRange(0, b[y])
			s := fa1.queryRange(0, b[y])

			now -= x*B[y] + asum - 2*((x-c)*B[y]+s)
			bsum -= B[y]
			fb0.update(b[y], -1)
			fb1.update(b[y], -B[y])
		}
		ans[q.id] = int64(now)
	}

	return ans
}

type fenwick []int

func (tr fenwick) update(i int, v int) {
	i++
	for i < len(tr) {
		tr[i] += v
		i += i & -i
	}
}

func (tr fenwick) query(i int) int {
	i++
	var res int
	for i > 0 {
		res += tr[i]
		i -= i & -i
	}
	return res
}

func (tr fenwick) queryRange(l, r int) int {
	return tr.query(r) - tr.query(l-1)
}
