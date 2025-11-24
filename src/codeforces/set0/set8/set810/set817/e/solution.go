package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, cur := range res {
		fmt.Fprintln(writer, cur)
	}
}

func drive(reader *bufio.Reader) []int {
	var q int
	fmt.Fscan(reader, &q)
	events := make([][]int, q)
	for i := 0; i < q; i++ {
		var t, p, l int
		fmt.Fscan(reader, &t, &p)
		if t == 3 {
			fmt.Fscan(reader, &l)
		}
		events[i] = []int{t, p, l}
	}
	return solve(events)
}

func solve(events [][]int) []int {
	tr := new(Trie)
	tr.next()

	var ans []int

	for _, cur := range events {
		switch cur[0] {
		case 1:
			tr.Update(cur[1], 1)
		case 2:
			tr.Update(cur[1], -1)
		default:
			var node int
			var sum int
			p, l := cur[1], cur[2]

			for i := H - 1; i >= 0; i-- {
				c := (l >> i) & 1
				b := (p >> i) & 1
				if c == 1 {
					if tr.children[node][b] != 0 {
						sum += tr.cnt[tr.children[node][b]]
					}
					node = tr.children[node][b^1]
				} else {
					node = tr.children[node][b]
				}

				if node == 0 {
					break
				}
			}
			ans = append(ans, sum)
		}
	}

	return ans
}

const H = 30

type Trie struct {
	children [][2]int
	cnt      []int
}

func (tr *Trie) next() int {
	tr.children = append(tr.children, [2]int{0, 0})
	tr.cnt = append(tr.cnt, 0)
	return len(tr.children) - 1
}

func (tr *Trie) Update(a int, v int) {
	var cur int
	for i := H - 1; i >= 0; i-- {
		x := (a >> i) & 1
		if tr.children[cur][x] == 0 {
			tr.children[cur][x] = tr.next()
		}
		cur = tr.children[cur][x]
		tr.cnt[cur] += v
	}
}
