package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, "YES")
	for _, v := range res {
		fmt.Fprintln(writer, v)
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

func drive(reader *bufio.Reader) (cnt []int, requests []string, res []string) {
	cnt = readNNums(reader, 6)
	n := readNum(reader)
	requests = make([]string, n)
	for i := range n {
		requests[i] = readString(reader)
	}
	res = solve(cnt, requests)
	return
}

type person struct {
	id           int
	firstChoice  int
	secondChoice int
}

var sizes = []string{"S", "M", "L", "XL", "XXL", "XXXL"}

func findIndex(s string) int {
	return slices.Index(sizes, s)
}

func parseRequest(id int, s string) person {
	ww := strings.Split(s, ",")
	firstChoice := findIndex(ww[0])
	secondChoice := -1
	if len(ww) > 1 {
		secondChoice = findIndex(ww[1])
	}
	return person{id, firstChoice, secondChoice}
}

func solve(cnt []int, requests []string) []string {
	n := len(requests)
	arr := make([]person, n)
	for i := range n {
		arr[i] = parseRequest(i, requests[i])
	}

	slices.SortFunc(arr, func(a, b person) int {
		return cmp.Or(a.firstChoice-b.firstChoice, a.secondChoice-b.secondChoice, a.id-b.id)
	})

	ans := make([]string, n)

	var pos int
	for i, v := range cnt {
		for v > 0 && pos < n && (arr[pos].firstChoice == i || arr[pos].secondChoice == i) {
			ans[arr[pos].id] = sizes[i]
			pos++
			v--
		}
		if pos < n && arr[pos].firstChoice == i && arr[pos].secondChoice == -1 {
			// 当前这个人只愿意使用这个size
			return nil
		}
	}

	if pos < n {
		return nil
	}

	return ans
}
