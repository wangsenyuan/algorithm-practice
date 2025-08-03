package p3636

import (
	"math"
	"slices"
)

func subarrayMajority(nums []int, queries [][]int) []int {
	freq := make(map[int]int)
	n := len(nums)

	set := NewSegTree(n + 1)

	trs := make([]*Node, n+1)

	change := func(num int, d int) {
		u := freq[num]
		if u > 0 {
			trs[u] = Delete(trs[u], num)
			if trs[u].Size() == 0 {
				set.Update(u, -1)
			}
		}
		u += d
		freq[num] = u
		if u > 0 {
			trs[u] = Insert(trs[u], num)
			set.Update(u, u)
		}
	}

	type query struct {
		id int
		l  int
		r  int
		th int
	}

	qs := make([]query, len(queries))
	for i, cur := range queries {
		qs[i] = query{
			id: i,
			l:  cur[0],
			r:  cur[1],
			th: cur[2],
		}
	}

	blockSize := int(math.Sqrt(float64(n))) + 1

	slices.SortFunc(qs, func(a, b query) int {
		if a.r/blockSize != b.r/blockSize {
			return a.r - b.r
		}
		d := a.r / blockSize
		if d&1 == 0 {
			return a.l - b.l
		}
		return b.l - a.l
	})

	var l, r int

	ans := make([]int, len(queries))

	for _, q := range qs {
		for r <= q.r {
			change(nums[r], 1)
			r++
		}
		// r > q.r
		for l > q.l {
			l--
			change(nums[l], 1)
		}
		// l == q.l
		for r > q.r+1 {
			r--
			change(nums[r], -1)
		}
		for l < q.l {
			change(nums[l], -1)
			l++
		}
		mf := set.Get(0, n+1)
		if mf < q.th {
			ans[q.id] = -1
		} else {
			ans[q.id] = trs[mf].MinValue()
		}
	}

	return ans
}

const inf = 1 << 60

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = -1
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v

	for p > 1 {
		tr[p>>1] = max(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int = -1
	for l < r {
		if l&1 == 1 {
			res = max(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	sz          int
	sum         int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.sz = 1
	node.sum = key
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	//y.height = max(y.left.Height(), y.right.Height()) + 1
	y.update()
	//x.height = max(x.left.Height(), x.right.Height()) + 1
	x.update()

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	//x.height = max(x.left.Height(), x.right.Height()) + 1
	//y.height = max(y.left.Height(), y.right.Height()) + 1
	x.update()
	y.update()

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.sz
}

func (node *Node) Sum() int {
	if node == nil {
		return 0
	}
	return node.sum
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.sz = node.left.Size() + node.right.Size() + node.cnt
	node.sum = node.left.Sum() + node.right.Sum() + node.cnt*node.key
}

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		node.cnt++
		node.sum += key
		node.sz++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	//node.height = max(node.left.Height(), node.right.Height()) + 1
	node.update()

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
func (root *Node) MaxValue() int {
	node := root
	for node.right != nil {
		node = node.right
	}
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
		root.cnt--
		root.sz--
		root.sum -= key
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

			root.key = tmp.key
			root.cnt = tmp.cnt
			root.sum = tmp.sum
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return root
	}

	//root.height = max(root.left.Height(), root.right.Height()) + 1
	root.update()

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
