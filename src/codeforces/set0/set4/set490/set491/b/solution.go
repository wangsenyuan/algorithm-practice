package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, optimalDistance, restaurantIndex := drive(reader)
	fmt.Println(optimalDistance)
	fmt.Println(restaurantIndex)
}

func drive(reader *bufio.Reader) (hotels [][]int, restaurants [][]int, optimalDistance int, restaurantIndex int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	var h int
	fmt.Fscan(reader, &h)
	hotels = make([][]int, h)
	for i := range h {
		hotels[i] = make([]int, 2)
		fmt.Fscan(reader, &hotels[i][0], &hotels[i][1])
	}
	var c int
	fmt.Fscan(reader, &c)
	restaurants = make([][]int, c)
	for i := range c {
		restaurants[i] = make([]int, 2)
		fmt.Fscan(reader, &restaurants[i][0], &restaurants[i][1])
	}
	optimalDistance, restaurantIndex = solve(n, m, hotels, restaurants)
	return
}

func solve(n int, m int, hotels [][]int, restaurants [][]int) (optimalDistance int, restaurantIndex int) {

	// 如果从上往下扫过来，那么左半边的到当前位置的距离 = cur[0] - tmp[0] + cur[1] - tmp[1] = cur[0] + cur[1] - (tmp[0] + tmp[1])
	// 也就是要找到最小的 tmp[0] + tmp[1] (因为要找到最大值)
	// 如果是右边部分 = cur[0] - tmp[0] + tmp[1] - cur[1] = (cur[0] - cur[1]) - (tmp[0] - tmp[1])
	// 也就是要找到tmp[0] - tmp[1]的最小值
	var cols []int
	for _, cur := range hotels {
		cols = append(cols, cur[1])
	}
	for _, cur := range restaurants {
		cols = append(cols, cur[1])
	}

	slices.Sort(cols)
	cols = slices.Compact(cols)
	t1 := NewSegTree(len(cols))
	t2 := NewSegTree(len(cols))

	ids := make([]int, len(restaurants))
	for i := range ids {
		ids[i] = i
	}

	slices.SortFunc(ids, func(i, j int) int {
		a := restaurants[i]
		b := restaurants[j]
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	var hotelId int

	slices.SortFunc(hotels, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	optimalDistance = inf
	dist := make([]int, len(restaurants))

	for _, i := range ids {
		cur := restaurants[i]
		for hotelId < len(hotels) && hotels[hotelId][0] <= cur[0] {
			hotel := hotels[hotelId]
			j := sort.SearchInts(cols, hotel[1])
			w := t1.Get(j, j+1)
			if w > hotel[0]+hotel[1] {
				t1.Update(j, hotel[0]+hotel[1])
			}
			w = t2.Get(j, j+1)
			if w > hotel[0]-hotel[1] {
				t2.Update(j, hotel[0]-hotel[1])
			}
			hotelId++
		}
		j := sort.SearchInts(cols, cur[1])
		dist[i] = max(cur[0]+cur[1]-t1.Get(0, j+1), cur[0]-cur[1]-t2.Get(j, len(cols)))
	}

	t1.Reset()
	t2.Reset()

	slices.SortFunc(hotels, func(a, b []int) int {
		return cmp.Or(b[0]-a[0], a[1]-b[1])
	})

	slices.SortFunc(ids, func(i int, j int) int {
		a := restaurants[i]
		b := restaurants[j]
		return cmp.Or(b[0]-a[0], a[1]-b[1])
	})

	// 当从下往上的时候， 左边部分 = tmp[0] - cur[0] + cur[1] - tmp[1] = (tmp[0] - tmp[1]) - (cur[0] - cur[1])
	// 所以要找 tmp[1] - tmp[0]的最小值
	// 如果是右边的话 tmp[0] - cur[0] + tmp[1] - cur[1] = (tmp[0] + tmp[1]) - (cur[0] + cur[1])
	hotelId = 0
	for _, i := range ids {
		cur := restaurants[i]
		for hotelId < len(hotels) && hotels[hotelId][0] >= cur[0] {
			hotel := hotels[hotelId]
			j := sort.SearchInts(cols, hotel[1])
			w := t1.Get(j, j+1)
			if w > hotel[1]-hotel[0] {
				t1.Update(j, hotel[1]-hotel[0])
			}
			w = t2.Get(j, j+1)
			if w > -(hotel[0] + hotel[1]) {
				t2.Update(j, -(hotel[0] + hotel[1]))
			}
			hotelId++
		}
		j := sort.SearchInts(cols, cur[1])
		dist[i] = max(dist[i], -t1.Get(0, j+1)-(cur[0]-cur[1]), -t2.Get(j, len(cols))-(cur[0]+cur[1]))
	}

	optimalDistance = slices.Min(dist)

	for i := range dist {
		if dist[i] == optimalDistance {
			restaurantIndex = i + 1
			break
		}
	}
	return
}

const inf = 1 << 60

func abs(a int) int {
	return max(a, -a)
}

func dist(a []int, b []int) int {
	dx := abs(a[0] - b[0])
	dy := abs(a[1] - b[1])
	return dx + dy
}

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = inf
	}
	return SegTree(arr)
}

func (t SegTree) Update(p int, v int) {
	p += len(t) / 2
	t[p] = v
	for p > 1 {
		t[p>>1] = min(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegTree) Get(l int, r int) int {
	var res int = inf
	l += len(t) / 2
	r += len(t) / 2
	for l < r {
		if l&1 == 1 {
			res = min(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func (t SegTree) Reset() {
	for i := range t {
		t[i] = inf
	}
}
