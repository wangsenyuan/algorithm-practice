package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		s := readString(reader)
		ss := strings.Split(s, " ")
		l, r := ss[0], ss[1]
		res := solve(l, r)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const inf = 1 << 60

func solve(l string, r string) int {
	n := len(l)
	dp := make([][]int, n+1)
	dp[0] = []int{inf, inf, inf, 0}

	for i := 0; i < n; i++ {
		dp[i+1] = []int{inf, inf, inf, inf}
		for state := range 4 {
			if dp[i][state] == inf {
				continue
			}
			sl := state / 2
			sr := state % 2
			u := 0
			if sl == 1 {
				u = int(l[i] - '0')
			}
			v := 9
			if sr == 1 {
				v = int(r[i] - '0')
			}
			for w := u; w <= v; w++ {
				nl := sl
				if sl == 0 || w > u {
					nl = 0
				}
				nr := sr
				if sr == 0 || w < v {
					nr = 0
				}
				dp[i+1][nl*2+nr] = min(dp[i+1][nl*2+nr], dp[i][state]+check(w == int(l[i]-'0'))+check(w == int(r[i]-'0')))
			}
		}
	}
	return slices.Min(dp[n])
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}
