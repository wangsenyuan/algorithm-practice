package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, _, res := drive(reader)
	fmt.Println(res[0], res[1], res[2], res[3])
}

func drive(reader *bufio.Reader) (t int, tp int, tu int, td int, a [][]int, res []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m, &t, &tp, &tu, &td)
	a = make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
	}
	for i := range n {
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	res = solve(a, t, tp, tu, td)
	return
}

const inf = 1 << 30

func solve(a [][]int, t int, tp int, tu int, td int) []int {
	n := len(a)
	m := len(a[0])
	// 固定两行r1, r2, 在这两行中间，迭代c2, 找到最优的c1
	// 不符合队列的性质，所以还得用到tree
	// 假设 pref[j]表示从0移动到j的时间
	// 对于c2来说, pref1[c2] - pref1[c1] + pref2[c2] - pref2[c1] + f(c2) + f(c1) >= t
	// f(c1)  + pref1[c1] + pref2[c1]<= sum + f(c2) - t 的最大值
	// 或者是  是 t - (sum + f(c2))的最小值
	up_down := make([][]int, n)
	down_up := make([][]int, n)
	for i := range n {
		up_down[i] = make([]int, m)
		down_up[i] = make([]int, m)
	}

	move := func(a int, b int) int {
		if a < b {
			return tu
		}
		if a > b {
			return td
		}
		return tp
	}

	for i := 1; i < n; i++ {
		for j := range m {
			up_down[i][j] = up_down[i-1][j] + move(a[i-1][j], a[i][j])
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := range m {
			down_up[i][j] = down_up[i+1][j] + move(a[i+1][j], a[i][j])
		}
	}

	ans := []int{0, 0, 0, 0, inf}

	update := func(r1 int, c1 int, r2 int, c2 int, val int) {
		diff := abs(val - t)
		if diff < ans[4] {
			ans[0] = r1 + 1
			ans[1] = c1 + 1
			ans[2] = r2 + 1
			ans[3] = c2 + 1
			ans[4] = diff
		}
	}
	f1 := make([]int, m)
	f2 := make([]int, m)
outer:
	for r1 := range n {
		clear(f1)
		for j := 1; j < m; j++ {
			f1[j] = f1[j-1] + move(a[r1][j-1], a[r1][j])
		}
		// 中间至少要有一行

		for r2 := r1 + 2; r2 < n; r2++ {
			c2 := m - 1
			var tr *Node
			clear(f2)
			for c1 := m - 2; c1 >= 0; c1-- {
				f2[c1] = f2[c1+1] + move(a[r2][c1+1], a[r2][c1])
				for c2-1 > c1 {
					tmp := up_down[r2][c2] - up_down[r1][c2] + f1[c2] - f2[c2]
					tr = Insert(tr, Pair{first: tmp, second: c2})
					c2--
				}
				if tr != nil {
					// 这个地方得放过来计算, 还需要知道位置
					tmp := down_up[r1][c1] - down_up[r2][c1] - f1[c1] + f2[c1]
					// tmp1 + tmp2 <= t
					node := LowerBound(tr, t-tmp)
					if node != nil {
						update(r1, c1, r2, node.item.second, tmp+node.item.first)
					}
					// tmp1 + tmp2 >= t
					node = UpperBound(tr, t-tmp)
					if node != nil {
						update(r1, c1, r2, node.item.second, tmp+node.item.first)
					}
				}
			}
			if ans[4] == 0 {
				break outer
			}
		}
	}

	return ans[:4]
}

func abs(num int) int {
	return max(num, -num)
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first || this.first == that.first && this.second < that.second
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Pair
	height      int
	size        int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.size
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func NewNode(item Pair) *Node {
	node := new(Node)
	node.item = item
	node.height = 1
	node.size = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + 1
	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + 1
	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + 1

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, item Pair) *Node {
	if node == nil {
		return NewNode(item)
	}
	if node.item == item {
		return node
	}
	if item.Less(node.item) {
		node.left = Insert(node.left, item)
	} else {
		node.right = Insert(node.right, item)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + 1
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

func Delete(root *Node, item Pair) *Node {
	if root == nil {
		return nil
	}

	if item.Less(root.item) {
		root.left = Delete(root.left, item)
	} else if root.item.Less(item) {
		root.right = Delete(root.right, item)
	} else {
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.item = tmp.item
			// make sure tmp node deleted after call delete on root.right
			root.right = Delete(root.right, tmp.item)
		}
	}

	if root == nil {
		return root
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	root.size = root.left.Size() + root.right.Size() + 1

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

func LowerBound(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.item.first <= key {
		res := LowerBound(node.right, key)
		if res != nil {
			return res
		}
		return node
	}
	return LowerBound(node.left, key)
}

func UpperBound(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.item.first >= key {
		res := UpperBound(node.left, key)
		if res != nil {
			return res
		}
		return node
	}
	return UpperBound(node.right, key)
}
