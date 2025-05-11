package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	people := make([][]int, n)
	for i := range n {
		people[i] = readNNums(reader, 2)
	}
	return solve(people)
}

type pair struct {
	first  int
	second int
}

func solve(people [][]int) []int {
	n := len(people)
	s := min(int(math.Sqrt(float64(n)))+1, 300)
	m := (n + s - 1) / s
	lists := make([]*DbList, m)
	for i := range m {
		lists[i] = NewDbList()
	}

	rebuild := func() {
		for j := 0; j < m; j++ {
			tmp := lists[j].PopBack(s)
			if len(tmp) == 0 {
				continue
			}
			for i := len(tmp) - 1; i >= 0; i-- {
				lists[j+1].AddFront(tmp[i])
			}
		}
	}

	for i := 0; i < n; i++ {
		a, c := people[i][0], people[i][1]
		j := i / s
		for c > 0 && j > 0 {
			if lists[j].sz > c || lists[j].mv > a {
				break
			}
			c -= lists[j].sz
			j--
		}

		// 在lists[j]中交换c次
		tmp := lists[j].tail
		for c > 0 && tmp != nil && tmp.val.first < a {
			tmp = tmp.prev
			c--
		}
		if tmp == nil {
			// 到头了
			lists[j].AddFront(pair{a, i})
		} else {
			lists[j].AddAfter(tmp, pair{a, i})
		}
		if lists[j].sz > 2*s {
			rebuild()
		}
	}

	res := make([]int, n)

	for i, j := 0, 0; i < m; i++ {
		for tmp := lists[i].head; tmp != nil; tmp = tmp.next {
			res[j] = tmp.val.second + 1
			j++
		}
	}

	return res
}

const inf = 1 << 60

type Node struct {
	prev *Node
	next *Node
	val  pair
}

func NewNode(val pair) *Node {
	node := new(Node)
	node.val = val
	return node
}

type DbList struct {
	sz   int
	mv   int
	head *Node
	tail *Node
}

func NewDbList() *DbList {
	list := new(DbList)
	list.sz = 0
	list.mv = -inf
	list.head = nil
	list.tail = nil
	return list
}

func (list *DbList) AddBack(v pair) {
	list.sz++
	list.mv = max(list.mv, v.first)
	node := NewNode(v)
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		node.prev = list.tail
		list.tail = node
	}
}

func (list *DbList) AddFront(v pair) {
	list.sz++
	list.mv = max(list.mv, v.first)
	node := NewNode(v)
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.head.prev = node
		node.next = list.head
		list.head = node
	}
}

func (list *DbList) AddAfter(node *Node, v pair) {
	list.sz++
	list.mv = max(list.mv, v.first)
	tmp := NewNode(v)
	tmp.next = node.next
	if node.next != nil {
		node.next.prev = tmp
	}
	node.next = tmp
	tmp.prev = node
	if list.tail == node {
		list.tail = tmp
	}
}

func (list *DbList) PopBack(k int) []pair {
	if list.sz <= k {
		return nil
	}

	list.sz = 0
	list.mv = -inf
	tmp := list.head
	for list.sz < k {
		list.mv = max(list.mv, tmp.val.first)
		list.sz++
		tmp = tmp.next
	}
	list.tail = tmp.prev
	list.tail.next = nil
	var res []pair

	for tmp != nil {
		res = append(res, tmp.val)
		tmp = tmp.next
	}

	return res
}
