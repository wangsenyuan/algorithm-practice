package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 998244353

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, 6)
		for j := range 6 {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(n, k, c, a)
}

func solve(n, k int, c []int, a [][]int) int {
	x := make([]int64, n)
	y := make([]int64, n)
	for i := 0; i < n; i++ {
		var sum, sum2 int64
		for j := 0; j < 6; j++ {
			v := int64(a[i][j])
			sum += v
			sum2 += v * v
		}
		x[i] = sum
		y[i] = 6*sum2 - sum*sum - 36*int64(c[i])
	}

	order := make([]int, n)
	for i := range n {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		u, v := order[i], order[j]
		if x[u] != x[v] {
			return x[u] < x[v]
		}
		if y[u] != y[v] {
			return y[u] > y[v]
		}
		return u < v
	})

	pos := make([]int, n)
	for i, id := range order {
		pos[id] = i
	}

	var sumX, sumY int64
	for i := 0; i < k; i++ {
		id := order[i]
		sumX += x[id]
		sumY += y[id]
	}
	best := sumX*sumX + sumY

	events := make([]Event, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			den := x[i] - x[j]
			if den == 0 {
				continue
			}
			num := y[j] - y[i]
			if den < 0 {
				den = -den
				num = -num
			}
			g := gcd(abs(num), den)
			events = append(events, Event{i, j, num / g, den / g})
		}
	}
	sort.Slice(events, func(i, j int) bool {
		a, b := events[i], events[j]
		return a.num*b.den < b.num*a.den
	})

	for l := 0; l < len(events); {
		r := l + 1
		for r < len(events) && sameEventTime(events[l], events[r]) {
			r++
		}

		if r == l+1 {
			u, v := events[l].u, events[l].v
			lo, hi := pos[u], pos[v]
			if lo > hi {
				lo, hi = hi, lo
			}
			reverseSegment(order, pos, lo, hi, k, x, y, &sumX, &sumY)
			best = max(best, sumX*sumX+sumY)
			l = r
			continue
		}

		id := make(map[int]int)
		for i := l; i < r; i++ {
			if _, ok := id[events[i].u]; !ok {
				id[events[i].u] = len(id)
			}
			if _, ok := id[events[i].v]; !ok {
				id[events[i].v] = len(id)
			}
		}
		uf := NewUFSet(len(id))
		for i := l; i < r; i++ {
			uf.Union(id[events[i].u], id[events[i].v])
		}

		groups := make(map[int][]int)
		for original, compressed := range id {
			root := uf.Find(compressed)
			groups[root] = append(groups[root], original)
		}

		segments := make([]Pair, 0, len(groups))
		for _, ids := range groups {
			if len(ids) < 2 {
				continue
			}
			lo, hi := n, -1
			for _, v := range ids {
				lo = min(lo, pos[v])
				hi = max(hi, pos[v])
			}
			segments = append(segments, Pair{lo, hi})
		}
		sort.Slice(segments, func(i, j int) bool {
			return segments[i].first < segments[j].first
		})

		for _, seg := range segments {
			reverseSegment(order, pos, seg.first, seg.second, k, x, y, &sumX, &sumY)
		}
		best = max(best, sumX*sumX+sumY)
		l = r
	}

	ans := best % mod
	if ans < 0 {
		ans += mod
	}
	ans = ans * pow(36, mod-2) % mod
	return int(ans)
}

type Event struct {
	u, v     int
	num, den int64
}

type Pair struct {
	first, second int
}

func sameEventTime(a, b Event) bool {
	return a.num == b.num && a.den == b.den
}

func reverseSegment(order []int, pos []int, l, r, k int, x, y []int64, sumX, sumY *int64) {
	if l < k && k <= r {
		cnt := k - l
		for i := l; i < k; i++ {
			id := order[i]
			*sumX -= x[id]
			*sumY -= y[id]
		}
		for i := r - cnt + 1; i <= r; i++ {
			id := order[i]
			*sumX += x[id]
			*sumY += y[id]
		}
	}
	for l < r {
		order[l], order[r] = order[r], order[l]
		pos[order[l]] = l
		pos[order[r]] = r
		l++
		r--
	}
}

type UFSet struct {
	arr []int
	cnt []int
}

func NewUFSet(n int) *UFSet {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := range n {
		arr[i] = i
		cnt[i] = 1
	}
	return &UFSet{arr, cnt}
}

func (uf *UFSet) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UFSet) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.arr[pb] = pa
	uf.cnt[pa] += uf.cnt[pb]
	return true
}

func pow(a, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
