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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
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

func drive(reader *bufio.Reader) []int {
	n := readNum(reader)
	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		queries[i] = make([]int, 2)
		s := readString(reader)
		var x int
		readInt([]byte(s), 2, &x)
		switch s[0] {
		case '+':
			queries[i] = []int{1, x}
		case '-':
			queries[i] = []int{2, x}
		default:
			queries[i] = []int{3, x}
		}
	}
	return solve(queries)
}

func solve(queries [][]int) []int {
	tr := new(Trie)
	tr.next()

	tr.Update(0, 1)

	var ans []int

	for _, cur := range queries {
		x := cur[1]
		switch cur[0] {
		case 1:
			tr.Update(x, 1)
		case 2:
			tr.Update(x, -1)
		default:
			y := tr.Find(x)
			ans = append(ans, y)
		}
	}

	return ans
}

const H = 30

type Trie struct {
	nodes [][2]int
	cnt   []int
}

func (tr *Trie) next() int {
	tr.nodes = append(tr.nodes, [2]int{})
	tr.cnt = append(tr.cnt, 0)
	return len(tr.nodes) - 1
}

func (tr *Trie) Update(num int, v int) {
	var node int
	for i := H - 1; i >= 0; i-- {
		b := (num >> i) & 1
		if tr.nodes[node][b] == 0 {
			tr.nodes[node][b] = tr.next()
		}
		node = tr.nodes[node][b]
		tr.cnt[node] += v
	}
}

func (tr *Trie) Find(x int) int {
	var y int
	var node int
	for i := H - 1; i >= 0; i-- {
		b := (x >> i) & 1
		c := tr.nodes[node][b^1]
		if c != 0 && tr.cnt[c] > 0 {
			y |= 1 << i
		} else {
			c = tr.nodes[node][b]
		}
		node = c
	}
	return y
}
