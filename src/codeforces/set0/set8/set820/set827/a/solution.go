package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	n := readNum(reader)
	requirements := make([]string, n)
	for i := range n {
		requirements[i] = readString(reader)
	}
	return solve(requirements)
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

type data struct {
	id    int
	s_len int
	pos   int
}

func solve(requirements []string) string {
	n := len(requirements)
	s := make([]string, n)
	var arr []data
	for i, cur := range requirements {
		var j int
		for j < len(cur) && cur[j] != ' ' {
			j++
		}
		s[i] = cur[:j]
		var k int
		j = readInt([]byte(cur), j+1, &k) + 1
		for range k {
			var p int
			j = readInt([]byte(cur), j, &p) + 1
			arr = append(arr, data{i, len(s[i]), p - 1})
		}
	}

	// 按照开始位置排序，相同位置的，按照长度排列
	slices.SortFunc(arr, func(a, b data) int {
		return cmp.Or(a.pos-b.pos, b.s_len-a.s_len)
	})
	var final_len int
	// m := len(arr)
	for _, cur := range arr {
		final_len = max(final_len, cur.pos+cur.s_len)
	}
	buf := make([]byte, final_len)
	for i := range final_len {
		buf[i] = 'a'
	}
	var pos int
	for _, cur := range arr {
		if cur.pos >= pos {
			copy(buf[cur.pos:], s[cur.id])
			pos = cur.pos + cur.s_len
		} else if pos < cur.pos+cur.s_len {
			// cur.pos < pos
			// pos - cur.pos:
			copy(buf[pos:], s[cur.id][pos-cur.pos:])
			pos = cur.pos + cur.s_len
		}
	}

	return string(buf)
}
