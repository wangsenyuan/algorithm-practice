package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	s := readString(reader)
	t := readString(reader)
	return solve(n, s, t)
}

const inf = 1 << 60

func solve(n int, s string, t string) int {
	pos := make([][]int, 26)
	m := len(s)
	for i := 0; i < m; i++ {
		x := int(s[i] - 'a')
		pos[x] = append(pos[x], i)
	}

	tot := m * n

	check := func(k int) bool {
		if k == 0 {
			return true
		}
		var j int
		for i := 0; i < len(t); i++ {
			if j >= tot {
				return false
			}
			x := int(t[i] - 'a')
			if len(pos[x]) == 0 {
				// 没有这个字符
				return false
			}
			// j1是真实的位置
			j1 := j % m
			// l是j1后面第一个x的位置
			l := search(len(pos[x]), func(l int) bool {
				return pos[x][l] >= j1
			})
			w := len(pos[x]) - l
			if w >= k {
				// 后面有足够的x,要移动到后面一个位置去
				r := pos[x][l+k-1] + 1
				j += r - j1
				continue
			}
			// w < k
			// 先把w消耗掉，j必须从一个新的S的开始位置
			need := k - w
			j += m - j1
			u, v := need/len(pos[x]), need%len(pos[x])

			if v > 0 {
				j += u * m
				j += pos[x][v-1]
				j++
			} else {
				j += (u - 1) * m
				j += pos[x][len(pos[x])-1] + 1
			}
		}

		return j <= tot
	}

	return search(inf, func(k int) bool {
		return !check(k)
	}) - 1
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
