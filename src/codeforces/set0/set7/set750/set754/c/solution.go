package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	tc := readNum(reader)
	for range tc {
		ans := drive(reader)
		if len(ans) == 0 {
			fmt.Fprintln(writer, "Impossible")
		} else {
			for _, cur := range ans {
				fmt.Fprintln(writer, cur)
			}
		}
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}
func drive(reader *bufio.Reader) []string {
	readNum(reader)
	users := readString(reader)
	m := readNum(reader)
	messages := make([]string, m)
	for i := range m {
		messages[i] = readString(reader)
	}
	return solve(users, messages)
}

func solve(users string, messages []string) []string {
	uu := strings.Split(users, " ")

	mem := make(map[string]int)
	for i, u := range uu {
		mem[u] = i
	}

	n := len(uu)
	m := len(messages)

	items := make([]*Item, m)

	findUserId := func(name string) int {
		if i, found := mem[name]; found {
			return i
		}
		return -1
	}

	var all [2]uint64
	for i := range n {
		if i >= 64 {
			all[1] |= 1 << (i - 64)
		} else {
			all[0] |= 1 << i
		}
	}

	contents := make([]string, m)

	mention := make([][2]uint64, m)

	var pq PriorityQueue
	sender := make([]int, m)

	for i := range m {
		it := new(Item)
		it.id = i

		sep := strings.Index(messages[i], ":")
		name := messages[i][:sep]
		contents[i] = messages[i][sep+1:]

		for pos := 0; pos < len(contents[i]); {
			if !isLetter(contents[i][pos]) {
				pos++
				continue
			}
			j := pos
			for pos < len(contents[i]) && isLetter(contents[i][pos]) {
				pos++
			}
			word := contents[i][j:pos]
			userId := findUserId(word)
			if userId >= 64 {
				mention[i][1] |= 1 << (userId - 64)
			} else if userId >= 0 {
				mention[i][0] |= 1 << userId
			}
		}

		userId := findUserId(name)
		if userId >= 0 {
			if userId >= 64 {
				it.value2 = 1 << (userId - 64)
			} else {
				it.value1 = 1 << userId
			}
		} else {
			it.value1 = all[0] ^ mention[i][0]
			it.value2 = all[1] ^ mention[i][1]
		}

		if it.priority() == 0 {
			return nil
		}

		sender[i] = userId
		heap.Push(&pq, it)
		items[i] = it
	}

	remove := func(id int, userId int) {
		if items[id].index < 0 {
			return
		}
		if userId >= 64 && items[id].value2&(1<<(userId-64)) != 0 {
			items[id].value2 ^= (1 << (userId - 64))
		}
		if userId < 64 && items[id].value1&(1<<userId) != 0 {
			items[id].value1 ^= (1 << userId)
		}
		heap.Fix(&pq, items[id].index)
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.priority() == 0 {
			// no anwer
			return nil
		}
		id := it.id
		if sender[id] < 0 {
			for i := range n {
				if i >= 64 && it.value2&(1<<(i-64)) != 0 {
					sender[id] = i
				}
				if i < 64 && it.value1&(1<<i) != 0 {
					sender[id] = i
				}
				if sender[id] >= 0 {
					break
				}
			}
		}

		if id > 0 {
			remove(id-1, sender[id])
		}
		if id < m-1 {
			remove(id+1, sender[id])
		}
	}

	ans := make([]string, m)

	for i := range m {
		u := uu[sender[i]]
		ans[i] = u + ":" + contents[i]
	}

	return ans
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
}

type Item struct {
	id     int
	value1 uint64
	value2 uint64
	index  int
}

func (it *Item) priority() int {
	x := bits.OnesCount64(it.value1)
	y := bits.OnesCount64(it.value2)
	return x + y
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority() < pq[j].priority() || pq[i].priority() == pq[j].priority() && pq[i].id < pq[j].id
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
	it := old[n-1]
	*pq = old[:n-1]
	it.index = -1
	return it
}
