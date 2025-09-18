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
	s := strings.Join(res, "\n")
	fmt.Print(s)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}
func drive(reader *bufio.Reader) []string {
	n := readNum(reader)
	res := make([]string, n)
	for i := range n {
		s := readString(reader)
		res[i] = solve(s)
	}
	return res
}

func solve(s string) string {
	n := len(s)
	if n&1 == 1 {
		return strings.Repeat("9", n-1)
	}

	freq := make([]int, 10)

	check := func(i int) bool {
		// 在保证 t < s 的情况下，能否得到最大的回文
		// 有这么多个位置可以使用
		var odd int
		for _, v := range freq {
			if v&1 == 1 {
				odd++
			}
		}

		return odd <= n-i
	}

	pos := -1
	val := -1
	for i := 0; i < n; i++ {
		// 如果在这里放置一个比较小的数
		var x int
		if i == 0 {
			x++
		}
		y := int(s[i] - '0')
		for j := y - 1; j >= x; j-- {
			freq[j]++
			if check(i + 1) {
				freq[j]--
				pos = i
				val = j
				break
			}
			freq[j]--
		}
		freq[y]++
	}

	if pos == -1 {
		return strings.Repeat("9", n-2)
	}
	clear(freq)
	for i := range pos {
		x := int(s[i] - '0')
		freq[x]++
	}
	freq[val]++
	buf := []byte(s)
	buf[pos] = byte(val + '0')
	// 要把这些位置给分配掉
	cnt := n - pos - 1
	assign := make([]int, 10)

	// 然后把奇数个的分配掉
	for x, v := range freq {
		if v&1 == 1 {
			assign[x]++
			cnt--
		}
	}
	// 剩下的都分配给9
	assign[9] += cnt
	for i, x := pos+1, 9; i < n; i++ {
		for assign[x] == 0 {
			x--
		}
		buf[i] = byte(x + '0')
		assign[x]--
	}

	return string(buf)
}
