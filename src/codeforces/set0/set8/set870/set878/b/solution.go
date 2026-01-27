package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k, m int
	fmt.Fscan(reader, &n, &k, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, m, a)
}

type data struct {
	cnt int
	val int
	pos int
}

func solve(k int, m int, a []int) int {
	var arr []data

	for i := range len(a) {
		if len(arr) > 0 && arr[len(arr)-1].val == a[i] {
			arr[len(arr)-1].cnt++
		} else {
			arr = append(arr, data{1, a[i], i})
		}
		if arr[len(arr)-1].cnt == k {
			arr = arr[:len(arr)-1]
		}
	}

	n := len(arr)

	var sum int
	for i := range n {
		sum += arr[i].cnt
	}

	var rem int

	var i int
	for i < n {
		if arr[i].val != arr[n-1-i].val {
			break
		}
		cnt := arr[i].cnt + arr[n-1-i].cnt
		rem += cnt / k * k
		cnt %= k
		if cnt > 0 {
			break
		}
		i++
	}

	if i == n {
		if m%2 == 1 {
			return sum
		}
		return 0
	}

	if 2*i+1 == n {
		s := m * arr[i].cnt
		s %= k
		if s > 0 {
			return sum - arr[i].cnt + s
		}
		return 0
	}

	return sum*m - rem*(m-1)
}
