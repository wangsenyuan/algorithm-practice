package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res, _, _, _, _ := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(writer, "NO")
		return
	}
	fmt.Fprintln(writer, "YES")
	for _, v := range res {
		fmt.Fprint(writer, len(v))
		for _, w := range v {
			fmt.Fprint(writer, " ", w)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) (res [][]int, n int, k int, p int, a []int) {
	fmt.Fscan(reader, &n, &k, &p)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a, k, p)
	return
}

func solve(a []int, k int, p int) [][]int {
	n := len(a)
	var odd []int
	var even []int
	for i := range n {
		if a[i]&1 == 1 {
			odd = append(odd, a[i])
		} else {
			even = append(even, a[i])
		}
	}
	if len(odd) < k-p {
		return nil
	}
	var res [][]int
	for i := range k - p {
		res = append(res, []int{odd[i]})
	}
	odd = odd[k-p:]

	// 每个上面都必须是偶数个奇数
	if len(odd)%2 == 1 {
		return nil
	}

	w := min(p, len(even))
	for i := range w {
		res = append(res, []int{even[i]})
	}
	even = even[w:]

	if len(res) < k {
		if len(odd)/2+len(res) < k {
			return nil
		}
		w = min(k-len(res), len(odd)/2)
		for i := range w {
			res = append(res, []int{odd[i*2], odd[i*2+1]})
		}
		odd = odd[2*w:]
	}

	if len(even) > 0 {
		if p > 0 {
			// 最后一个上面全部是偶数
			res[k-1] = append(res[k-1], even...)
		} else {
			// p == 0
			res[0] = append(res[0], even...)
		}
	}

	if len(odd) > 0 {

		res[0] = append(res[0], odd...)
	}

	return res
}
