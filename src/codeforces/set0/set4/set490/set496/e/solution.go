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
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (parts [][]int, actors [][]int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	parts = make([][]int, n)
	for i := range parts {
		parts[i] = make([]int, 2)
		fmt.Fscan(reader, &parts[i][0], &parts[i][1])
	}
	var m int
	fmt.Fscan(reader, &m)
	actors = make([][]int, m)
	for i := range actors {
		actors[i] = make([]int, 3)
		fmt.Fscan(reader, &actors[i][0], &actors[i][1], &actors[i][2])
	}
	res = solve(parts, actors)
	return parts, actors, res
}

func solve(parts [][]int, actors [][]int) []int {
	var nums []int
	for _, cur := range parts {
		nums = append(nums, cur[0])
	}
	for _, cur := range actors {
		nums = append(nums, cur[0])
	}
	sort.Ints(nums)
	nums = slices.Compact(nums)

	k := len(nums)
	a1 := make([][]int, k)
	a2 := make([][]int, k)

	for i, cur := range parts {
		j := sort.SearchInts(nums, cur[0])
		a1[j] = append(a1[j], i)
	}
	cnt := make([]int, len(actors))
	for i, cur := range actors {
		j := sort.SearchInts(nums, cur[0])
		a2[j] = append(a2[j], i)
		cnt[i] = cur[2]
	}

	var tr *Node

	m := len(parts)
	ans := make([]int, m)
	for i := range nums {
		for _, j := range a2[i] {
			d := actors[j][1]
			tr = Insert(tr, Pair{d, j})
		}

		for _, j := range a1[i] {
			b := parts[j][1]
			tmp := FindEqualOrGreater(tr, Pair{b, -1})
			if tmp == nil {
				return nil
			}
			v := tmp.item.second
			ans[j] = v + 1
			cnt[v]--
			if cnt[v] == 0 {
				tr = Delete(tr, tmp.item)
			}
		}
	}

	return ans
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	if this.first < that.first {
		return true
	}
	if this.first == that.first && (this.second < that.second) {
		return true
	}
	return false
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Pair
	height      int32
	cnt         int32
	left, right *Node
}

func (node *Node) Height() int32 {
	if node == nil {
		return 0
	}
	return node.height
}

func NewNode(item Pair) *Node {
	node := new(Node)
	node.item = item
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

func (node *Node) GetBalance() int32 {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func FindEqualOrGreater(node *Node, item Pair) *Node {
	if node == nil {
		return nil
	}
	if node.item == item {
		return node
	}

	if node.item.Less(item) {
		return FindEqualOrGreater(node.right, item)
	}

	res := FindEqualOrGreater(node.left, item)
	if res == nil {
		return node
	}
	return res
}

func FindEqualOrLess(node *Node, item Pair) *Node {
	if node == nil {
		return nil
	}
	if item.Less(node.item) {
		return FindEqualOrLess(node.left, item)
	}

	res := FindEqualOrLess(node.right, item)
	if res == nil {
		return node
	}
	return res
}

func Insert(node *Node, item Pair) *Node {
	if node == nil {
		return NewNode(item)
	}

	if node.item == item {
		node.cnt++
		return node
	}

	if item.Less(node.item) {
		node.left = Insert(node.left, item)
	} else {
		node.right = Insert(node.right, item)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	balance := node.GetBalance()

	if balance > 1 && item.Less(node.left.item) {
		return rightRotate(node)
	}

	if balance < -1 && node.right.item.Less(item) {
		return leftRotate(node)
	}

	if balance > 1 && node.left.item.Less(item) {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && item.Less(node.right.item) {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func MinValueNode(root *Node) *Node {
	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func MaxValueNode(root *Node) *Node {
	cur := root

	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

func Delete(root *Node, item Pair) *Node {
	if root == nil {
		return nil
	}

	if item.Less(root.item) {
		root.left = Delete(root.left, item)
	} else if root.item.Less(item) {
		root.right = Delete(root.right, item)
	} else {
		root.cnt--
		if root.cnt > 0 {
			return root
		}

		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.item = tmp.item
			root.cnt = tmp.cnt
			tmp.cnt = 1

			// make sure tmp node deleted after call delete on root.right
			root.right = Delete(root.right, tmp.item)
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
