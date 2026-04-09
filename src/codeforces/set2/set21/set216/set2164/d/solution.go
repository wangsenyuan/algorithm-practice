package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	tc := readNums(reader)[0]
	for range tc {
		_, _, ok, res := drive(reader)
		if !ok {
			fmt.Fprintln(writer, -1)
			continue
		}
		fmt.Fprintln(writer, len(res))
		for _, s := range res {
			fmt.Fprintln(writer, s)
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i := range res {
		res[i], _ = strconv.Atoi(ss[i])
	}
	return res
}

func drive(reader *bufio.Reader) (s string, t string, ok bool, res []string) {
	nums := readNums(reader)
	k := nums[1]
	s = readString(reader)
	t = readString(reader)
	ok, res = solve(s, t, k)
	return
}

func solve(s string, t string, kmx int) (ok bool, res []string) {
	if s[0] != t[0] {
		return false, nil
	}
	n := len(s)
	pos := make([][]int, 26)
	for i := range n {
		x := int(s[i] - 'a')
		pos[x] = append(pos[x], i)
	}

	buf := []byte(s)
	todo := make([][]int, kmx+1)

	var k int

	for i := n - 1; i > 0; i-- {
		x := int(buf[i] - 'a')
		pos[x] = pos[x][:len(pos[x])-1]
		if buf[i] == t[i] {
			continue
		}
		// buf[i] != t[i]
		w := int(t[i] - 'a')
		if len(pos[w]) == 0 {
			return false, nil
		}
		i1 := pos[w][len(pos[w])-1]
		if i-i1 > kmx {
			return false, nil
		}
		k = max(k, i-i1)

		for j := i - 1; j > i1; j-- {
			x := int(buf[j] - 'a')
			pos[x] = pos[x][:len(pos[x])-1]
		}

		for j := i1 + 1; j <= i; j++ {
			todo[j-i1] = append(todo[j-i1], j)
			buf[j] = buf[j-1]
			x := int(buf[j] - 'a')
			if j < i {
				pos[x] = append(pos[x], j)
			}
		}
	}

	buf = []byte(s)
	for i := 1; i <= k; i++ {
		// todo is in descending order
		for _, j := range todo[i] {
			buf[j] = buf[j-1]
		}
		res = append(res, string(buf))
	}

	return true, res
}
