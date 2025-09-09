package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) []int {
	n := len(a)
	assign := make([]int, 256)
	for i := range 256 {
		assign[i] = -1
	}

	res := make([]int, n)

	assign[0] = 0

	freq := make([]int, 256)
	freq[0]++

	for i := range n {
		if assign[a[i]] == -1 {
			j := a[i]
			for assign[j] == -1 {
				j--
			}
			cnt := a[i] - j
			if freq[assign[j]]+cnt <= k {
				for z := j + 1; z <= a[i]; z++ {
					assign[z] = assign[j]
					freq[assign[z]]++
				}
			} else {
				// 这里希望能得到 x := a[i] - k + 1 这个值，
				// 但是中间剩余 (x - assign[j]) * k >= x - j 否则的话，中间会出现空的set
				x := a[i]
				for x-1 >= a[i]-k+1 && x-1 > j && (x-1-assign[j]+1)*k >= x-1-j {
					x--
				}
				for z := x; z <= a[i]; z++ {
					assign[z] = x
					freq[x]++
				}
			}
		}
		res[i] = assign[a[i]]
	}

	return res
}
