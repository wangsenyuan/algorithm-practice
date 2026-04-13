package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

// maxA: problem bound on a[i]. maxHyp: hypotenuse can exceed maxA when both legs are ≤ maxA.
const maxA = 10_000_000
const maxHyp = 15_000_000

// posScratch[value] = lawn index + 1, or 0 if absent. One buffer avoids reallocating ~40 MiB per run/test.
var posScratch []int32

func ensurePosScratch() {
	if posScratch == nil {
		posScratch = make([]int32, maxA+1)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func solve(a []int) int {
	n := len(a)
	ensurePosScratch()
	pos := posScratch
	for i := range a {
		pos[a[i]] = int32(i + 1)
	}
	defer func() {
		for i := range a {
			pos[a[i]] = 0
		}
	}()

	set := NewDSU(n)
	enumeratePrimitiveTriples(pos, set)

	res := 0
	for i := range n {
		if set.Find(i) == i {
			res++
		}
	}
	return res
}

func enumeratePrimitiveTriples(pos []int32, set *DSU) {
	uniteVals := func(u, v int) {
		pu := pos[u]
		pv := pos[v]
		if pu != 0 && pv != 0 {
			set.Union(int(pu-1), int(pv-1))
		}
	}

	for m := 2; m*m+1 <= maxHyp; m++ {
		for k := 1; k < m && m*m+k*k <= maxHyp; k++ {
			if gcd(m, k) != 1 || (m^k)&1 == 0 {
				continue
			}
			legA := m*m - k*k
			legB := 2 * m * k
			hyp := m*m + k*k
			l1, l2 := legA, legB
			if l1 > l2 {
				l1, l2 = l2, l1
			}
			if l1 <= 0 || l1 > maxA || l2 > maxA {
				continue
			}
			if gcd(l1, l2) != 1 || gcd(l1, hyp) != 1 || gcd(l2, hyp) != 1 {
				continue
			}
			uniteVals(l1, l2)
			if hyp <= maxA {
				uniteVals(l1, hyp)
				uniteVals(l2, hyp)
			}
		}
	}
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := range n {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (d *DSU) Find(x int) int {
	if d.arr[x] != x {
		d.arr[x] = d.Find(d.arr[x])
	}
	return d.arr[x]
}

func (d *DSU) Union(x, y int) bool {
	px := d.Find(x)
	py := d.Find(y)
	if px == py {
		return false
	}
	if d.cnt[px] < d.cnt[py] {
		px, py = py, px
	}
	d.cnt[px] += d.cnt[py]
	d.arr[py] = px
	return true
}
