package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)

	for _, x := range res {
		fmt.Fprintln(writer, x[0], x[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	frogs := make([][]int, n)
	for i := range n {
		frogs[i] = make([]int, 2)
		fmt.Fscan(reader, &frogs[i][0], &frogs[i][1])
	}

	mosquitos := make([][]int, m)
	for i := range m {
		mosquitos[i] = make([]int, 2)
		fmt.Fscan(reader, &mosquitos[i][0], &mosquitos[i][1])
	}

	return solve(frogs, mosquitos)
}

type frog struct {
	id int
	x  int
	t  int
}

func solve(frogs [][]int, mosquitos [][]int) [][]int {
	n := len(frogs)
	arr := make([]frog, n)
	ans := make([][]int, n)

	R := make([]int, n)

	for i := range n {
		ans[i] = make([]int, 2)
		ans[i][1] = frogs[i][1]
		arr[i] = frog{
			id: i,
			x:  frogs[i][0],
			t:  frogs[i][1],
		}
		R[i] = arr[i].x + arr[i].t
	}

	slices.SortFunc(arr, func(a, b frog) int {
		return a.x - b.x
	})

	xs := make([]int, n)
	reach := make([]int, n)
	for i := range n {
		xs[i] = arr[i].x
		reach[i] = arr[i].x + arr[i].t
	}
	seg := NewSegTree(reach)

	var wait *Node

	play := func(p int, b int) {
		pos := sort.SearchInts(xs, p+1) - 1
		id := -1
		if pos >= 0 {
			id = seg.FirstAtLeast(pos, p)
		}
		if id < 0 {
			wait = Insert(wait, p, b)
			return
		}

		orig := arr[id].id
		reach[id] += b
		ans[orig][0]++
		ans[orig][1] += b
		seg.Update(id, reach[id])

		for {
			tmp := FindPrevNode(wait, reach[id])
			if tmp == nil || tmp.key < xs[id] {
				break
			}
			ans[orig][0] += tmp.cnt
			ans[orig][1] += tmp.val
			reach[id] += tmp.val
			wait = Delete(wait, tmp.key)
			seg.Update(id, reach[id])
		}
	}

	for _, cur := range mosquitos {
		play(cur[0], cur[1])
	}

	return ans
}

type SegTree struct {
	n   int
	arr []int
}

func NewSegTree(vals []int) *SegTree {
	n := 1
	for n < len(vals) {
		n <<= 1
	}
	arr := make([]int, 2*n)
	for i, v := range vals {
		arr[n+i] = v
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = max(arr[i<<1], arr[i<<1|1])
	}
	return &SegTree{n, arr}
}

func (tr *SegTree) Update(pos int, val int) {
	pos += tr.n
	tr.arr[pos] = val
	for pos > 1 {
		pos >>= 1
		tr.arr[pos] = max(tr.arr[pos<<1], tr.arr[pos<<1|1])
	}
}

func (tr *SegTree) FirstAtLeast(r int, val int) int {
	return tr.firstAtLeast(1, 0, tr.n-1, r, val)
}

func (tr *SegTree) firstAtLeast(node int, l int, r int, qr int, val int) int {
	if l > qr || tr.arr[node] < val {
		return -1
	}
	if l == r {
		return l
	}
	mid := (l + r) / 2
	res := tr.firstAtLeast(node<<1, l, mid, qr, val)
	if res >= 0 {
		return res
	}
	return tr.firstAtLeast(node<<1|1, mid+1, r, qr, val)
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	val         int
	height      int
	cnt         int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func NewNode(key int, val int) *Node {
	node := new(Node)
	node.key = key
	node.val = val
	node.height = 1
	node.cnt = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	x.height = max(x.left.Height(), x.right.Height()) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	y.height = max(y.left.Height(), y.right.Height()) + 1

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, key int, val int) *Node {
	if node == nil {
		return NewNode(key, val)
	}
	if node.key == key {
		node.val += val
		node.cnt++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key, val)
	} else {
		node.right = Insert(node.right, key, val)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	balance := node.GetBalance()

	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func MinValueNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func (root *Node) MinValue() int {
	if root == nil {
		return 0
	}
	node := MinValueNode(root)

	return node.key
}

func Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = Delete(root.left, key)
	} else if key > root.key {
		root.right = Delete(root.right, key)
	} else {
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.key = tmp.key
			root.cnt = tmp.cnt
			root.val = tmp.val
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return root
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	balance := root.GetBalance()

	if balance > 1 && root.left.GetBalance() >= 0 {
		return rightRotate(root)
	}

	if balance > 1 && root.left.GetBalance() < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	if balance < -1 && root.right.GetBalance() <= 0 {
		return leftRotate(root)
	}

	if balance < -1 && root.right.GetBalance() > 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}

// find max node.key <= key
func FindPrevNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key <= key {
		res := FindPrevNode(node.right, key)
		if res == nil {
			res = node
		}
		return res
	}
	return FindPrevNode(node.left, key)
}
