package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
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

func drive(reader *bufio.Reader) []int {
	readString(reader)
	s := readString(reader)
	m := readNum(reader)
	queries := make([]string, m)
	for i := range m {
		queries[i] = readString(reader)
	}
	return solve(s, queries)
}

type data struct {
	id int
	m  int
}

func solve(s string, queries []string) []int {
	n := len(s)

	qs := make([][]data, 26)
	for i, cur := range queries {
		var m int
		pos := readInt([]byte(cur), 0, &m)
		c := int(cur[pos+1] - 'a')
		qs[c] = append(qs[c], data{i, m})
	}

	sum := make([]int, n+1)
	get := func(c int, m int) int {
		var j int
		var best int
		for i := range n {
			x := int(s[i] - 'a')
			sum[i+1] = sum[i]
			if x != c {
				sum[i+1]++
			}
			for j < i && sum[i+1]-sum[j] > m {
				j++
			}
			best = max(best, i+1-j)
		}

		return best
	}

	ans := make([]int, len(queries))

	for c, vs := range qs {
		slices.SortFunc(vs, func(a, b data) int {
			return a.m - b.m
		})

		for i := 0; i < len(vs); {
			j := i
			for i < len(vs) && vs[i].m == vs[j].m {
				i++
			}
			res := get(c, vs[j].m)
			for j < i {
				ans[vs[j].id] = res
				j++
			}
		}
	}

	return ans
}
