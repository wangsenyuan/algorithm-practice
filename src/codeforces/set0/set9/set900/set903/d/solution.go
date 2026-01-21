package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/big"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res.String())
}

func drive(reader *bufio.Reader) big.Int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) big.Int {
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(x pair, y pair) int {
		return cmp.Or(x.first-y.first, x.second-y.second)
	})

	// 这个符号要根据 (i, j)的关系决定
	// 貌似会溢出

	sum := make(BIT, n+2)
	cnt := make(BIT, n+2)

	res := big.NewInt(0)

	update := func(a int, b int) {
		// add a - b
		a -= b
		w := big.NewInt(int64(a))
		res = res.Add(res, w)
	}

	for l, r := 0, 0; r < n; r++ {

		for l < n && arr[l].first+1 < arr[r].first {
			sum.update(arr[l].second, arr[l].first)
			cnt.update(arr[l].second, 1)
			l++
		}

		i := arr[r].second
		// 它前面的取正值，后面的取负值
		s1 := sum.query(0, i)
		c1 := cnt.query(0, i)
		update(c1*arr[r].first, s1)
		s2 := sum.query(i+1, n)
		c2 := cnt.query(i+1, n)
		update(s2, c2*arr[r].first)
	}

	return *res
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}
