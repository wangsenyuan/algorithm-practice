package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		var buf bytes.Buffer
		for _, cur := range res {
			buf.WriteString(cur)
			buf.WriteByte('\n')
		}
		buf.WriteTo(os.Stdout)
	}
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

func process(reader *bufio.Reader) ([]string, []string) {
	n := readNum(reader)
	persons := make([]string, n)
	for i := 0; i < n; i++ {
		persons[i] = readString(reader)
	}
	return solve(persons), persons
}

type Person struct {
	id   int
	name string
	a    int
	h    int
}

const X = 30000

func solve(persons []string) []string {
	n := len(persons)
	arr := make([]Person, n)
	for i, cur := range persons {
		var j int
		for cur[j] != ' ' {
			j++
		}
		name := cur[:j]
		a, _ := strconv.Atoi(cur[j+1:])
		arr[i] = Person{
			id:   i,
			name: name,
			a:    a,
		}
	}

	slices.SortFunc(arr, func(a, b Person) int {
		return a.a - b.a
	})

	ans := make([]Person, n+1)

	cur_h := X

	for i, cur := range arr {
		a := cur.a
		if i < a {
			return nil
		}
		// i >= a
		if a == 0 {
			// 他排在哪里都没有关系
			cur.h = X
			ans[i] = cur
		} else {
			cur.h = cur_h
			copy(ans[a+1:], ans[a:])
			ans[a] = cur
		}
		cur_h--
	}

	res := make([]string, n)
	for i := range n {
		res[i] = fmt.Sprintf("%s %d", ans[i].name, ans[i].h)
	}

	return res
}
