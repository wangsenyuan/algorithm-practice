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

func drive(reader *bufio.Reader) (ok bool, p []int) {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	return solve(n, k)
}

func solve(n, k int) (ok bool, p []int) {
	k1 := k
	k ^= n
	res := make([]int, n)
	marked := make([]bool, n)
	pos := n
	for i := n - 1; i > 0; i-- {
		if k < n {
			pos = i
			marked[k] = true
			res[i] = k
			break
		}
		// k >= n
		h := bits.Len(uint(k)) - 1
		w := 1 << h
		if w >= n {
			return false, nil
		}
		k -= w
		res[i] = w
		marked[1<<h] = true
	}
	if !marked[0] {
		// 保证0最后一个出现, 这样可以不用考虑mex出现多次的情况
		res[pos-1] = 0
		pos--
		marked[0] = true
	}

	for i, v := 0, 0; i < pos; i++ {
		if v >= n {
			return false, nil
		}
		for v < n && marked[v] {
			v++
		}
		if v == n {
			return false, nil
		}
		res[i] = v
		v++
	}
	var k2 int
	var mex int

	clear(marked)

	for i := range n {
		marked[res[i]] = true
		for mex < n && marked[mex] {
			mex++
		}
		if mex > n {
			return false, nil
		}
		k2 ^= mex
	}

	return k2 == k1, res
}
