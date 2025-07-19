package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	res := process(reader)
	for _, ans := range res {
		buf.WriteString(ans)
		buf.WriteByte('\n')
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []string {
	_, m, q := readThreeNums(reader)
	s := readString(reader)
	words := strings.Split(s, " ")
	relations := make([]string, m)
	for i := 0; i < m; i++ {
		relations[i] = readString(reader)
	}
	queries := make([]string, q)
	for i := 0; i < q; i++ {
		queries[i] = readString(reader)
	}
	return solve(words, relations, queries)
}

func solve(words []string, relations []string, queries []string) []string {
	ids := make(map[string]int)
	for i, word := range words {
		ids[word] = i
	}
	n := len(words)
	color := make([]int, n)
	gid := make([]int, n)
	for i := range n {
		color[i] = -1
		gid[i] = -1
	}

	var sets [][]int

	merge := func(x, y int, c int) {
		a := gid[x]
		b := gid[y]
		if len(sets[a]) > len(sets[b]) {
			a, b = b, a
			x, y = y, x
		}

		// len(sets[a]) <= len(sets[b])
		for _, v := range sets[a] {
			if v == x {
				// x不能变，否则就错掉了
				continue
			}
			gid[v] = b
			color[v] = color[y] ^ c ^ color[x] ^ color[v]
		}

		color[x] = c ^ color[y]
		gid[x] = b

		sets[b] = append(sets[b], sets[a]...)
		sets[a] = sets[a][:0]
	}

	var ans []string

	for _, relation := range relations {
		ss := strings.Split(relation, " ")
		x, y := ids[ss[1]], ids[ss[2]]

		if color[x] < 0 {
			color[x] = 0
			gid[x] = len(sets)
			sets = append(sets, []int{x})
		}
		if color[y] < 0 {
			color[y] = 0
			gid[y] = len(sets)
			sets = append(sets, []int{y})
		}

		if gid[x] == gid[y] {
			// 它们在同一组
			if (color[x] == color[y]) == (ss[0] == "1") {
				ans = append(ans, "YES")
			} else {
				ans = append(ans, "NO")
			}
		} else {
			merge(x, y, int(ss[0][0]-'1'))
			ans = append(ans, "YES")
		}
	}

	for _, query := range queries {
		ss := strings.Split(query, " ")
		x := ids[ss[0]]
		y := ids[ss[1]]
		if gid[x] != gid[y] || gid[x] < 0 {
			ans = append(ans, "3")
		} else {
			if color[x] == color[y] {
				ans = append(ans, "1")
			} else {
				ans = append(ans, "2")
			}
		}
	}

	return ans
}
