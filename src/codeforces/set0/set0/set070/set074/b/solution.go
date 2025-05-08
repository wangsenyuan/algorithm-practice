package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func process(reader *bufio.Reader) string {
	n, m, k := readThreeNums(reader)
	direction := readString(reader)
	train := readString(reader)
	return solve(n, m, k, direction, train)
}

type state struct {
	stowaway_pos         int
	controller_pos       int
	controller_direction int // 0 go head, 1 go tail
}

func solve(n int, m int, k int, direction string, train string) string {
	cur := state{
		stowaway_pos:         m,
		controller_pos:       k,
		controller_direction: 0,
	}
	if direction == "to tail" {
		cur.controller_direction = 1
	}

	for i := 0; i < len(train); i++ {
		next := cur
		if train[i] == '0' {
			// stowaway move
			if cur.controller_direction == 0 {
				// controller move to head
				if cur.stowaway_pos < cur.controller_pos {
					// 尽快向头部移动
					next.stowaway_pos = max(1, cur.stowaway_pos-1)
				} else {
					// next.stoway_pos > cur.controller_pos
					next.stowaway_pos = min(n, next.stowaway_pos+1)
				}
				next.controller_pos--
				if next.controller_pos == next.stowaway_pos {
					return fmt.Sprintf("Controller %d", i+1)
				}
				if next.controller_pos == 1 {
					next.controller_direction = 1
				}
			} else {
				if cur.stowaway_pos < cur.controller_pos {
					next.stowaway_pos = max(1, next.stowaway_pos-1)
				} else {
					next.stowaway_pos = min(n, next.stowaway_pos+1)
				}
				next.controller_pos++
				if next.controller_pos == next.stowaway_pos {
					return fmt.Sprintf("Controller %d", i+1)
				}
				if next.controller_pos == n {
					next.controller_direction = 0
				}
			}
		} else {
			if cur.controller_direction == 0 {
				next.stowaway_pos = n
				next.controller_pos--
				if next.controller_pos == 1 {
					next.controller_direction = 1
				}
			} else {
				next.stowaway_pos = 1
				next.controller_pos++
				if next.controller_pos == n {
					next.controller_direction = 0
				}
			}
		}
		cur = next
	}
	return "Stowaway"
}
