package p3569

const N = 1e5 + 10

func maximumCount(nums []int, queries [][]int) []int {
	var primes []int
	lpf := make([]int, N)
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= N {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
	m := len(primes)
	id := make([]int, N)
	for i, x := range primes {
		id[x] = i
	}

	pos := make([]*Node, m)
	for i, x := range nums {
		if lpf[x] == x {
			j := id[x]
			pos[j] = Insert(pos[j], i)
		}
	}
	n := len(nums)
	tr := NewTree(n)

	doIt := func(x int, s int) {
		u := MinValueNode(pos[id[x]])
		v := MaxValueNode(pos[id[x]])
		if u.key < v.key {
			tr.Update(u.key, v.key, s)
		}
	}
	var cnt int

	for j := range primes {
		if pos[j] != nil {
			cnt++
			doIt(primes[j], 1)
		}
	}

	update := func(i int, x int) {
		if nums[i] == x {
			return
		}
		y := nums[i]
		if lpf[y] == y {
			doIt(y, -1)
			pos[id[y]] = Delete(pos[id[y]], i)
			if pos[id[y]] != nil {
				doIt(y, 1)
			} else {
				cnt--
			}
		}
		nums[i] = x
		if lpf[x] == x {
			if pos[id[x]] != nil {
				doIt(x, -1)
			} else {
				cnt++
			}
			pos[id[x]] = Insert(pos[id[x]], i)
			doIt(x, 1)
		}
	}

	var res []int

	for _, cur := range queries {
		i, x := cur[0], cur[1]
		update(i, x)
		res = append(res, tr.Query(0, n-1)+cnt)
	}

	return res
}

type Tree struct {
	val  []int
	lazy []int
	n    int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &Tree{val, lazy, n}
}

func (t *Tree) update(i int, v int) {
	t.val[i] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.update(i*2+1, t.lazy[i])
		t.update(i*2+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.val[i] = max(t.val[i*2+1], t.val[i*2+2])
}

func (t *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			t.update(i, v)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		loop(i*2+1, l, mid)
		loop(i*2+2, mid+1, r)
		t.pull(i)
	}
	loop(0, 0, t.n-1)
}

func (t *Tree) Query(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) >> 1
		x := loop(i*2+1, l, mid)
		y := loop(i*2+2, mid+1, r)
		return max(x, y)
	}
	return loop(0, 0, t.n-1)
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
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

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
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

func Search(node *Node, key int) *Node {
	if node == nil || node.key == key {
		return node
	}
	if key < node.key {
		return Search(node.left, key)
	}
	return Search(node.right, key)
}

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		node.cnt++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
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
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
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
