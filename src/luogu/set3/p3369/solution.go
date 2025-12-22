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
	n := readNum(reader)
	cmds := make([][]int, n)
	for i := 0; i < n; i++ {
		cmds[i] = readNNums(reader, 2)
	}
	res := solve(cmds)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
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

func solve(cmds [][]int) []int {
	var root *Node

	nums := make(map[int]int)
	for _, cmd := range cmds {
		if cmd[0] == 4 {
			continue
		}
		nums[cmd[1]]++
	}
	arr := make([]int, 0, len(nums))
	for k := range nums {
		arr = append(arr, k)
	}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	for i := 0; i < len(arr); i++ {
		nums[arr[i]] = i
	}

	insert := func(x int) {
		var L, R *Node
		Split(root, x, &L, &R)
		cur := NewNode(x, nums[x])
		root = Merge(Merge(L, cur), R)
	}

	remove := func(x int) {
		var L, R, p *Node
		Split(root, x, &L, &R)
		Split(L, x-1, &L, &p)
		// remove x
		if p != nil {
			p = Merge(p.ls, p.rs)
		}
		root = Merge(Merge(L, p), R)
	}

	rank := func(x int) int {
		var L, R *Node
		Split(root, x-1, &L, &R)
		res := L.Size() + 1
		root = Merge(L, R)
		return res
	}

	precursor := func(x int) int {
		var L, R *Node
		Split(root, x-1, &L, &R)
		res := L.Kth(L.Size()).key
		root = Merge(L, R)
		return res
	}

	successor := func(x int) int {
		var L, R *Node
		Split(root, x, &L, &R)
		res := R.Kth(1).key
		root = Merge(L, R)
		return res
	}

	var res []int

	for _, cmd := range cmds {
		if cmd[0] == 1 {
			insert(cmd[1])
		} else if cmd[0] == 2 {
			remove(cmd[1])
		} else if cmd[0] == 3 {
			res = append(res, rank(cmd[1]))
		} else if cmd[0] == 4 {
			tmp := root.Kth(cmd[1])
			res = append(res, tmp.key)
		} else if cmd[0] == 5 {
			res = append(res, precursor(cmd[1]))
		} else {
			res = append(res, successor(cmd[1]))
		}
	}
	return res
}

type Node struct {
	ls, rs *Node
	key    int
	pri    int
	sz     int
}

func NewNode(k int, p int) *Node {
	node := new(Node)
	node.sz = 1
	node.key = k
	node.pri = p
	return node
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.sz
}

func (node *Node) update() {
	node.sz = node.ls.Size() + node.rs.Size() + 1
}

func Split(root *Node, x int, L **Node, R **Node) {
	if root == nil {
		*L = nil
		*R = nil
		return
	}
	if root.key <= x {
		*L = root
		Split(root.rs, x, &(root.rs), R)
	} else {
		*R = root
		Split(root.ls, x, L, &(root.ls))
	}
	root.update()
}

func Merge(L *Node, R *Node) *Node {
	if L == nil || R == nil {
		if R != nil {
			return R
		}
		return L
	}

	if L.pri > R.pri {
		L.rs = Merge(L.rs, R)
		L.update()
		return L
	}
	R.ls = Merge(L, R.ls)
	R.update()
	return R
}

func (node *Node) Kth(k int) *Node {
	if k == node.ls.Size()+1 {
		return node
	}
	if k <= node.ls.Size() {
		return node.ls.Kth(k)
	}
	return node.rs.Kth(k - node.ls.Size() - 1)
}
