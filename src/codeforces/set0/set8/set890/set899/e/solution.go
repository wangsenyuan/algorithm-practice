package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	nums := readNNums(reader, n)

	res := solve(nums)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

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

const inf = 1 << 60

func solve(nums []int) int {
	n := len(nums)

	next := make([]int, n)
	prev := make([]int, n)
	id := make([]int, n)

	for i := range n {
		prev[i] = i - 1
		id[i] = i
		if i > 0 && nums[i] == nums[i-1] {
			prev[i] = prev[i-1]
			id[i] = id[i-1]
		}
	}
	for i := n - 1; i >= 0; i-- {
		next[i] = i + 1
		if i+1 < n && nums[i] == nums[i+1] {
			next[i] = next[i+1]
		}
	}

	items := make([]*Item, n)
	var pq PriorityQueue
	for i := range n {
		if prev[i] == i-1 {
			it := new(Item)
			it.id = i
			it.val = next[i] - i
			items[i] = it
			heap.Push(&pq, it)
		}
	}

	var ans int
	for pq.Len() > 0 {
		ans++
		it := heap.Pop(&pq).(*Item)
		i := it.id
		l := prev[i]
		r := next[i]
		if l >= 0 && r < n && nums[l] == nums[r] {
			l = id[l]
			heap.Remove(&pq, items[r].index)
			items[l].val += items[r].val
			heap.Fix(&pq, items[l].index)

			next[l] = next[r]
			if next[r] < n {
				prev[next[r]] = l
			}
		} else {
			if l >= 0 {
				next[id[l]] = r
			}
			if r < n {
				prev[r] = l
			}
		}
	}

	return ans
}

func solve1(nums []int) int {
	n := len(nums)

	tr1 := NewSegTree(n, -1, max)
	tr2 := NewSegTree(n, n, min)

	for i := 0; i < n; i++ {
		tr1.Update(i, i)
		tr2.Update(i, i)
	}
	pq := make(PriorityQueue, 0, n)
	items := make([]*Item, n)

	sets := make([][]int, n)

	uf := NewUF(n)

	last := make([]int, n)
	for i := 0; i < n; {
		j := i
		var arr []int
		for i < n && nums[i] == nums[j] {
			uf.Union(j, i)
			arr = append(arr, i)
			i++
		}
		sets[j] = arr
		it := new(Item)
		it.id = j
		it.val = len(arr)
		items[j] = it
		heap.Push(&pq, it)
		last[j] = i - 1
	}

	var res int

	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*Item).id
		res++
		j, k := -1, n
		if i > 0 {
			// j是左边最近的没有被删除的位置
			j = tr1.Get(0, i)
		}

		if last[i]+1 < n {
			// k 是右边最近的没有被删除的位置
			k = tr2.Get(last[i]+1, n)
		}

		if j >= 0 {
			j = uf.Find(j)
		}

		if j >= 0 && k < n && nums[j] == nums[k] {
			// need to merge
			pq.update(items[k], inf)
			heap.Pop(&pq)
			a, b := sets[j], sets[k]
			if len(a) < len(b) {
				a, b = b, a
			}
			a = append(a, b...)
			sets[j] = a
			uf.Union(j, k)
			pq.update(items[j], len(a))
			last[j] = last[k]
		}

		for _, v := range sets[i] {
			// remove this range
			tr1.Update(v, -1)
			tr2.Update(v, n)
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type UF struct {
	arr []int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return &UF{arr}
}

func (uf *UF) Find(i int) int {
	if uf.arr[i] != i {
		uf.arr[i] = uf.Find(uf.arr[i])
	}
	return uf.arr[i]
}

func (uf *UF) Union(a, b int) {
	a = uf.Find(a)
	b = uf.Find(b)
	if a != b {
		if a > b {
			a, b = b, a
		}
		uf.arr[b] = a
	}
}

type Item struct {
	id    int
	val   int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].val > pq[j].val || pq[i].val == pq[j].val && pq[i].id < pq[j].id
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[0 : n-1]
	return res
}

func (pq *PriorityQueue) update(item *Item, value int) {
	item.val = value
	heap.Fix(pq, item.index)
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
