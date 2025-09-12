package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f %.10f\n", res[0], res[1])
}

func drive(reader *bufio.Reader) []float64 {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	w := n / k
	var q int
	fmt.Fscan(reader, &q)
	cards := make([][]int, q)
	for i := range q {
		cards[i] = make([]int, w)
		for j := range w {
			fmt.Fscan(reader, &cards[i][j])
		}
	}
	return solve(n, k, a, cards)
}

const inf = 1 << 60

func solve(n int, k int, a []int, cards [][]int) []float64 {
	marked := make([]bool, n+1)

	res := []float64{inf, 0}

	// w := n / k
	var cnt int
	for _, cur := range cards {
		if marked[cur[0]-1] {
			continue
		}
		cnt++
		var sum int
		for _, v := range cur {
			sum += a[v-1]
			marked[v-1] = true
		}
		avg := float64(sum) / float64(len(cur))
		res[0] = min(res[0], avg)
		res[1] = max(res[1], avg)
	}

	var arr []int
	for i := range n {
		if !marked[i] {
			arr = append(arr, a[i])
		}
	}
	sort.Ints(arr)

	// w := min(n/k, len(arr))
	w := n / k
	// 必须还得有一个card才行
	if cnt < k && len(arr) >= w {
		var sum int
		for i := range w {
			sum += arr[i]
		}
		res[0] = min(res[0], float64(sum)/float64(w))
		sum = 0
		for i := len(arr) - w; i < len(arr); i++ {
			sum += arr[i]
		}
		res[1] = max(res[1], float64(sum)/float64(w))
	}

	return res
}
