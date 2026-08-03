package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		ok, p := drive(reader)
		if !ok {
			fmt.Fprintln(writer, "NO")
			continue
		}
		fmt.Fprintln(writer, "YES")
		s := fmt.Sprintf("%v", p)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (bool, []int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a, b []int) (bool, []int) {
	n := len(a)
	for i := range n {
		a[i]--
		if b[i] > 0 {
			b[i]--
		}
	}

	vis := make([]bool, n)

	for i := range n {
		if b[i] != -1 && !vis[i] {
			j := i
			for !vis[j] {
				vis[j] = true
				if b[a[j]] == -1 {
					b[a[j]] = a[b[j]]
				} else if b[a[j]] != a[b[j]] {
					return false, nil
				}
				j = a[j]
			}
		}
	}

	cnt := make([]int, n)
	for i := range n {
		if b[i] != -1 {
			cnt[b[i]]++
		}
	}

	for i := range n {
		if cnt[i] > 1 {
			return false, nil
		}
	}

	vis1 := make([]bool, n)
	vis2 := make([]bool, n)
	cycl1 := make([][]int, n+1)
	cycl2 := make([][]int, n+1)
	for i := range n {
		if b[i] == -1 && !vis1[i] {
			j := i
			var cycleLen int
			for !vis1[j] {
				vis1[j] = true
				cycleLen++
				j = a[j]
			}
			cycl1[cycleLen] = append(cycl1[cycleLen], i)
		}
		if cnt[i] == 0 && !vis2[i] {
			j := i
			var cycleLen int
			for !vis2[j] {
				vis2[j] = true
				cycleLen++
				j = a[j]
			}
			cycl2[cycleLen] = append(cycl2[cycleLen], i)
		}
	}

	for k := 1; k <= n; k++ {
		for i := range cycl1[k] {
			x := cycl1[k][i]
			y := cycl2[k][i]
			for b[x] == -1 {
				b[x] = y
				x = a[x]
				y = a[y]
			}
		}
	}

	for i := range n {
		b[i]++
	}

	return true, b
}
