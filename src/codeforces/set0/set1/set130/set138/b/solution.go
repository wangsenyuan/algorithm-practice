package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, s1, s2 := process(reader)
	fmt.Println(s1)
	fmt.Println(s2)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (num string, s1 string, s2 string) {
	num = readString(reader)
	s1, s2 = solve(num)
	return
}

func solve(num string) (string, string) {
	var a [10]int
	for _, x := range []byte(num) {
		a[int(x-'0')]++
	}
	b := a

	var best int
	// 先找到一个配置，可以得到最大值的情况
	var ans int
	for x := 1; x < 10; x++ {
		if a[x] == 0 || b[10-x] == 0 {
			continue
		}
		c := a
		d := b
		c[x]--
		d[10-x]--
		cnt := 1
		for u := 0; u < 10; u++ {
			tmp := min(c[u], d[9-u])
			cnt += tmp
			c[u] -= tmp
			d[9-u] -= tmp
		}
		cnt += min(c[0], d[0])
		if cnt > best {
			best = cnt
			ans = x
		}
	}
	var buf1 []byte
	var buf2 []byte
	// n := len(num)
	if ans > 0 {
		buf1 = append(buf1, byte(ans+'0'))
		buf2 = append(buf2, byte(10-ans+'0'))
		a[ans]--
		b[10-ans]--
		for u := 0; u < 10; u++ {
			tmp := min(a[u], b[9-u])
			a[u] -= tmp
			b[9-u] -= tmp

			for tmp > 0 {
				buf1 = append(buf1, byte(u+'0'))
				buf2 = append(buf2, byte(9-u+'0'))
				tmp--
			}
		}
		buf1 = reverse(buf1)
		buf2 = reverse(buf2)
	}

	tmp := min(a[0], b[0])
	a[0] -= tmp
	b[0] -= tmp
	for tmp > 0 {
		buf1 = append(buf1, '0')
		buf2 = append(buf2, '0')
		tmp--
	}

	buf1 = reverse(buf1)
	buf2 = reverse(buf2)
	for x := 0; x < 10; x++ {
		for a[x] > 0 {
			buf1 = append(buf1, byte(x+'0'))
			a[x]--
		}
		for b[x] > 0 {
			buf2 = append(buf2, byte(x+'0'))
			b[x]--
		}
	}

	buf1 = reverse(buf1)
	buf2 = reverse(buf2)
	return string(buf1), string(buf2)
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
