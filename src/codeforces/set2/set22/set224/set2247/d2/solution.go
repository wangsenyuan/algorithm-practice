package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		ans := drive(reader)
		s := fmt.Sprint(ans)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	updates := make([][]int, q)
	for j := range q {
		var i, x int
		fmt.Fscan(reader, &i, &x)
		updates[j] = []int{i, x}
	}
	return solve(a, updates)
}

const inf = 1 << 60

func solve(a []int, updates [][]int) []int {
	n := len(a)
	h := bits.Len(uint(n))

	for len(a) < 1<<h {
		a = append(a, inf)
	}

	type node struct {
		minVal int
		maxVal int
		k      int
	}

	arr := make([]node, 4*len(a))

	merge := func(i int, l int, mid int) {
		arr[i].minVal = min(arr[i*2+1].minVal, arr[i*2+2].minVal)
		arr[i].maxVal = max(arr[i*2+1].maxVal, arr[i*2+2].maxVal)
		arr[i].k = max(arr[i*2+1].k, arr[i*2+2].k)
		if arr[i*2+1].maxVal > arr[i*2+2].minVal {
			arr[i].k = mid - l
		}
	}

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l+1 == r {
			arr[i] = node{a[l], a[l], 0}
		} else {
			mid := (l + r) >> 1
			build(i*2+1, l, mid)
			build(i*2+2, mid, r)
			merge(i, l, mid)
		}
	}

	build(0, 0, len(a))

	var play func(i int, l int, r int, pos int, v int)
	play = func(i int, l int, r int, pos int, v int) {
		if l+1 == r {
			arr[i].minVal = v
			arr[i].maxVal = v
			arr[i].k = 0
		} else {
			mid := (l + r) >> 1
			if pos < mid {
				play(i*2+1, l, mid, pos, v)
			} else {
				play(i*2+2, mid, r, pos, v)
			}
			merge(i, l, mid)
		}
	}

	ans := make([]int, len(updates)+1)
	ans[0] = arr[0].k

	for i, update := range updates {
		p, x := update[0], update[1]
		play(0, 0, len(a), p, x)
		ans[i+1] = arr[0].k
	}

	return ans
}
