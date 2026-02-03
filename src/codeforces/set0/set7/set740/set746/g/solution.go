package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, _, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, n)
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) (n int, k int, a []int, res [][]int) {
	var t int
	fmt.Fscan(reader, &n, &t, &k)

	a = make([]int, t)
	for i := range t {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(n, k, a)
	return
}

func solve(n int, k int, a []int) [][]int {
	// a[i] = 第i层的数量
	t := len(a)
	hi := make([]int, t)
	lo := make([]int, t)
	hi[t-1] = a[t-1]
	lo[t-1] = a[t-1]
	for i := t - 2; i >= 0; i-- {
		// 只有一个父节点，其他的都是叶子节点
		hi[i] = hi[i+1] + a[i] - 1
		// 尽量的消耗父节点
		lo[i] = lo[i+1] + max(0, a[i]-a[i+1])
	}

	if k < lo[0] || k > hi[0] {
		return nil
	}

	var res [][]int

	level := []int{1}
	id := 2
	var next []int
	for range a[0] {
		res = append(res, []int{level[0], id})
		next = append(next, id)
		id++
	}

	level = next

	for i := range t - 1 {
		// 假设这层需要y个叶子节点
		// k - y >= lo[i+1]
		// y <= k - lo[i+1] and y <= a[i] - 1
		// k - y >= hi[i+1] => y >= k - hi[i+1] 这个必须要成立
		y := min(a[i]-1, k-lo[i+1])
		// 那么剩余的x个节点，就可以作为父节点
		x := a[i] - y
		var next []int
		for j := range x {
			res = append(res, []int{level[j], id})
			next = append(next, id)
			id++
		}

		for j := x; j < a[i+1]; j++ {
			// 剩余的都将第一个作为父节点
			res = append(res, []int{level[0], id})
			next = append(next, id)
			id++
		}

		k -= y

		level = next
	}
	// k == a[t-1] holds

	return res
}
