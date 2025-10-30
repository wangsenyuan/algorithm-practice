package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	a := readNum(reader)
	s := readString(reader)
	return solve(a, s)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func solve(a int, s string) int {
	n := len(s)
	freq := make([]int, n*9+1)
	for i := range n {
		var sum int
		for j := i; j < n; j++ {
			sum += int(s[j] - '0')
			freq[sum]++
		}
	}
	var ans int

	for i := range n {
		var sum int
		for j := i; j < n; j++ {
			sum += int(s[j] - '0')
			// s2 * s1 = a
			// 如果 a = 0
			if sum == 0 {
				if a == 0 {
					// [i...j]行都是0， 且有n列
					ans += n * (n + 1) / 2
				}
				continue
			}
			// sum > 0
			if a%sum == 0 && a/sum < len(freq) {
				ans += freq[a/sum]
			}
		}
	}
	return ans
}
