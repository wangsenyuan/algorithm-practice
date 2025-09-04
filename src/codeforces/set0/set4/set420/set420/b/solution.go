package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	messages := make([]string, m)
	for i := range m {
		messages[i] = readString(reader)
	}
	return solve(n, messages)
}

func solve(n int, messages []string) []int {
	m := len(messages)

	ids := make([]int, m)

	last := make([]int, n+1)
	for i := range n + 1 {
		last[i] = -1
	}

	var arr []int

	for i, message := range messages {
		id, _ := strconv.Atoi(message[2:])
		ids[i] = id
		if message[0] == '-' {
			if last[id] < 0 {
				arr = append(arr, id)
			}
		}
		last[id] = i
	}

	leader := ids[0]
	if len(arr) > 0 {
		leader = arr[len(arr)-1]
	}
	// check leader ok or not

	level := max(0, len(arr)-1)

	ok := true
	occ := true

	for i := 0; i < m && ok; i++ {
		if ids[i] == leader {
			if level > 0 {
				ok = false
				break
			}
			if messages[i][0] == '-' {
				occ = false
			} else {
				occ = true
			}
		} else {
			if !occ {
				ok = false
				break
			}
			if messages[i][0] == '+' {
				level++
			} else {
				level--
			}
		}
	}

	var res []int
	if ok {
		res = append(res, leader)
	}

	for i := 1; i <= n; i++ {
		if last[i] < 0 {
			res = append(res, i)
		}
	}
	sort.Ints(res)

	return res
}
