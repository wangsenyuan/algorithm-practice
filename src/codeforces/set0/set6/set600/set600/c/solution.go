package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

const inf = 1 << 30

func solve(s string) string {
	freq := make([]int, 26)
	for i := range s {
		x := int(s[i] - 'a')
		freq[x]++
	}
	n := len(s)
	// 只需要修改奇数的那些
	var arr []int
	for i, v := range freq {
		if v&1 == 1 {
			arr = append(arr, i)
		}
	}
	// 要往两头去处理
	m := len(arr)
	// 修改的数量 = m / 2
	for l, r := 0, m-1; l < r; l, r = l+1, r-1 {
		freq[arr[l]]++
		freq[arr[r]]--
	}
	return play(n, freq)
}

func play(n int, freq []int) string {
	buf := make([]byte, n)
	var c int
	if n&1 == 1 {
		var mid int
		for i, v := range freq {
			if v&1 == 1 {
				mid = i
				break
			}
		}
		buf[n/2] = byte(mid + 'a')
		freq[mid]--
	}

	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		for freq[c] == 0 {
			c++
		}
		buf[l] = byte(c + 'a')
		freq[c]--
		buf[r] = byte(c + 'a')
		freq[c]--
	}

	return string(buf)
}
