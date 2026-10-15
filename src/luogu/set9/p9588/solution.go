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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var c, q int
	fmt.Fscan(reader, &c, &q)
	ops := make([][2]int, q)
	for i := range q {
		var op int
		fmt.Fscan(reader, &op)
		if op == 4 {
			ops[i] = [2]int{4, 0}
			continue
		}
		var v int
		fmt.Fscan(reader, &v)
		ops[i] = [2]int{op, v}
	}
	_ = c
	return solve(ops)
}

func solve(ops [][2]int) []int {
	m := len(ops)
	nums := []int{0}
	pref := []int64{0}
	st := make(segtree, 2*(m+1))
	var del int64
	gone := 0
	var ans []int

	for i := range ops {
		cur := ops[i]
		switch cur[0] {
		case 1:
			x := cur[1]
			nums = append(nums, x)
			pref = append(pref, pref[len(pref)-1]+int64(x))
			st.update(len(nums)-1, x)
		case 2:
			del += int64(cur[1])
			for gone+1 < len(nums) && pref[gone+1] <= del {
				gone++
				st.update(gone, 0)
			}
		case 3:
			pos := del + int64(cur[1])
			k := sort.Search(len(pref), func(i int) bool {
				return pref[i] >= pos
			})
			ans = append(ans, int(pos-pref[k-1]))
		default:
			ans = append(ans, st.get(1, len(nums)))
		}
	}

	return ans
}

type segtree []int

func (st segtree) update(p int, v int) {
	n := len(st) / 2
	p += n
	st[p] = v
	for p > 1 {
		st[p>>1] = max(st[p], st[p^1])
		p >>= 1
	}
}

func (st segtree) get(l int, r int) int {
	var res int
	n := len(st) / 2
	l += n
	r += n

	for l < r {
		if l&1 == 1 {
			res = max(res, st[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, st[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
