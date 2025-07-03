package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	diversity, ans := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", diversity))
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (int, []int) {
	s := readString(reader)
	return solve(s)
}

func solve(s string) (int, []int) {
	// 字符s[i] = x, 假设它上一个位置是 l[x], 那么左端点在l[x]...i中间，且右端点在i..n
	// 的所有子串的diversity +1
	// 那是不是l[x]...i这个区间内的（左端点组成的区间)+(n-i)
	n := len(s)

	marked := make([]int, 26)
	var diversity int
	for i := range n {
		x := int(s[i] - 'a')
		if marked[x] == 0 {
			diversity++
		}
		marked[x]++
	}

	marked2 := make([]int, 26)

	add := func(i int, marked []int, cnt int) int {
		x := int(s[i] - 'a')
		if marked[x] == 0 {
			cnt++
		}
		marked[x]++
		return cnt
	}

	rem := func(i int, marked []int, cnt int) int {
		x := int(s[i] - 'a')
		if marked[x] == 1 {
			cnt--
		}
		marked[x]--
		return cnt
	}

	check := func(expect int) int {
		clear(marked)
		clear(marked2)

		var res int
		var l1, cnt1 int
		var l2, cnt2 int
		for r := 0; r < n; r++ {
			cnt1 = add(r, marked, cnt1)
			cnt2 = add(r, marked2, cnt2)

			for l1 < r && cnt1 > expect {
				cnt1 = rem(l1, marked, cnt1)
				if cnt1 == expect {
					cnt1 = add(l1, marked, cnt1)
					break
				}
				l1++
			}

			for l2 < r && cnt2 >= expect {
				cnt2 = rem(l2, marked2, cnt2)
				if cnt2 < expect {
					cnt2 = add(l2, marked2, cnt2)
					break
				}
				l2++
			}
			if cnt2 == expect {
				if l1 == 0 && cnt1 == expect {
					res += l2 + 1
				} else {
					res += l2 - l1
				}
			}
		}

		return res
	}

	ans := make([]int, diversity)

	for x := range diversity {
		ans[x] = check(x + 1)
	}

	return diversity, ans
}
