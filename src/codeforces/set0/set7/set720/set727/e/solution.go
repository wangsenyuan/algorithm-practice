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
	_, _, res := process(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (string, []string, []int) {
	_, k := readTwoNums(reader)
	s := readString(reader)
	n := readNum(reader)
	g := make([]string, n)
	for i := 0; i < n; i++ {
		g[i] = readString(reader)
	}
	return s, g, solve(k, s, g)
}

func solve(k int, s string, g []string) []int {

	m := len(g)
	n := len(s)

	tr := NewTrie(m * k)

	for i, x := range g {
		tr.AddString(x, i+1)
	}

	dp := make([]int32, n+k+1)
	fp := make([]int32, n+k+1)
	for i := range len(dp) {
		dp[i] = -1
	}
	dp[0] = 0

	var state int32

	prev := make([]map[int]bool, m+1)
	for i := range m + 1 {
		prev[i] = make(map[int]bool)
	}

	for i := 0; i < n+k; i++ {
		state = tr.Transition(state, int32(s[i%n]-'a'))
		if tr.nodes[state].id != -1 {
			j := tr.nodes[state].id
			if prev[j][(i+1)%k] {
				continue
			}
			// 好像还是有问题
			dp[i+1] = j
			fp[i+1] = 1 + fp[i+1-k]
			prev[j][(i+1)%k] = true
		}
	}

	pos := -1
	for i := n; i <= n+k; i++ {
		if int(fp[i])*k == n {
			pos = i
			break
		}
	}

	if pos == -1 {
		return nil
	}

	var ans []int
	for i := pos; i > 0; i -= k {
		if dp[i] == -1 {
			break
		}
		ans = append(ans, int(dp[i]))
	}

	slices.Reverse(ans)

	return ans
}

const AL = 26

// const N = 5000
const INF = 1000000000

type Node struct {
	next       [AL]int32
	parent     int32
	parentByte int32
	suffixLink int32
	transition [AL]int32
	id         int32
}

func NewNode() *Node {
	node := new(Node)
	for i := 0; i < AL; i++ {
		node.next[i] = -1
		node.transition[i] = -1
	}
	node.suffixLink = -1
	node.parent = -1
	node.id = -1
	return node
}

type Trie struct {
	nodes []*Node
}

func NewTrie(sizeHint int) *Trie {
	trie := new(Trie)
	trie.nodes = make([]*Node, 0, sizeHint+1)
	trie.nodes = append(trie.nodes, NewNode())
	return trie
}

func (trie *Trie) AddString(s string, id int) {
	var v int32
	nodes := trie.nodes
	for i := 0; i < len(s); i++ {
		c := int32(s[i] - 'a')
		if nodes[v].next[c] == -1 {
			nodes[v].next[c] = int32(len(nodes))
			cur := NewNode()
			nodes = append(nodes, cur)
			cur.parent = v
			cur.parentByte = c
		}
		v = nodes[v].next[c]
	}
	nodes[v].id = int32(id)
	trie.nodes = nodes
}

func (trie *Trie) Transition(v int32, c int32) int32 {
	if trie.nodes[v].transition[c] == -1 {
		if trie.nodes[v].next[c] != -1 {
			trie.nodes[v].transition[c] = trie.nodes[v].next[c]
		} else {
			if v == 0 {
				trie.nodes[v].transition[c] = 0
			} else {
				trie.nodes[v].transition[c] = trie.Transition(trie.GetLink(v), c)
			}
		}
	}
	return trie.nodes[v].transition[c]
}

func (trie *Trie) GetLink(v int32) int32 {
	if trie.nodes[v].suffixLink == -1 {
		if v == 0 || trie.nodes[v].parent == 0 {
			trie.nodes[v].suffixLink = 0
		} else {
			trie.nodes[v].suffixLink =
				trie.Transition(trie.GetLink(trie.nodes[v].parent), trie.nodes[v].parentByte)
		}
	}

	return trie.nodes[v].suffixLink
}
