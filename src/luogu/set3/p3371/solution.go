package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	swaps := make([][]int, m)
	for i := 0; i < m; i++ {
		swaps[i] = readNNums(reader, 2)
	}
	res := solve(n, swaps)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}

	fmt.Println(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(n int, swaps [][]int) []int {
	var tr *Node

	for i := 1; i <= n; i++ {
		tr = Merge(tr, NewNode(i))
	}

	for _, cur := range swaps {
		x, y := cur[0], cur[1]
		var L, R, p *Node
		Split(tr, y, &L, &R)
		Split(L, x-1, &L, &p)
		p.lazy ^= 1
		tr = Merge(Merge(L, p), R)
	}

	var arr []int

	var inorder func(n *Node)
	inorder = func(n *Node) {
		if n == nil {
			return
		}
		n.pushDown()
		inorder(n.ls)
		arr = append(arr, n.num)
		inorder(n.rs)
	}
	inorder(tr)

	return arr
}

type Node struct {
	ls, rs *Node
	num    int
	pri    int
	size   int
	lazy   int
}

func NewNode(x int) *Node {
	n := new(Node)
	n.num = x
	n.size = 1
	n.pri = rand.Int()
	return n
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return n.size
}

func (n *Node) Update() {
	n.size = n.ls.Size() + n.rs.Size() + 1
}

func (n *Node) pushDown() {
	if n.lazy == 1 {
		n.ls, n.rs = n.rs, n.ls
		if n.ls != nil {
			n.ls.lazy ^= 1
		}
		if n.rs != nil {
			n.rs.lazy ^= 1
		}
		n.lazy = 0
	}
}

func Split(n *Node, x int, L, R **Node) {
	if n == nil {
		*L = nil
		*R = nil
		return
	}
	n.pushDown()
	if n.ls.Size()+1 <= x {
		*L = n
		Split(n.rs, x-n.ls.Size()-1, &n.rs, R)
	} else {
		*R = n
		Split(n.ls, x, L, &n.ls)
	}
	n.Update()
}

func Merge(L, R *Node) *Node {
	if L == nil || R == nil {
		if L != nil {
			return L
		}
		return R
	}
	if L.pri > R.pri {
		L.pushDown()
		L.rs = Merge(L.rs, R)
		L.Update()
		return L
	}
	R.pushDown()
	R.ls = Merge(L, R.ls)
	R.Update()
	return R
}
