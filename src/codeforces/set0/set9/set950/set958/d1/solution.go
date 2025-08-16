package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) string {
	m := readNum(reader)
	ships := make([]string, m)
	for i := range m {
		ships[i] = readString(reader)
	}
	res := solve(ships)
	s := fmt.Sprintf("%v", res)
	return s[1 : len(s)-1]
}

type ship struct {
	x int
	y int
}

func parse(s string) ship {
	// (a + b) / c
	parts := strings.Split(s, "/")
	c, _ := strconv.Atoi(parts[1])
	parts[0] = strings.Trim(parts[0], "()")
	parts = strings.Split(parts[0], "+")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])

	a += b
	g := gcd(a, c)

	return ship{a / g, c / g}
}

func solve(ships []string) []int {
	m := len(ships)
	arr := make([]ship, m)
	for i := range m {
		arr[i] = parse(ships[i])
	}
	sorted_arr := slices.Clone(arr)
	slices.SortFunc(sorted_arr, func(a ship, b ship) int {
		return cmp.Or(a.x-b.x, a.y-b.y)
	})

	ans := make([]int, m)

	for i := range m {
		r := sort.Search(m, func(j int) bool {
			return sorted_arr[j].x > arr[i].x || (sorted_arr[j].x == arr[i].x && sorted_arr[j].y > arr[i].y)
		})
		l := sort.Search(m, func(j int) bool {
			return sorted_arr[j].x > arr[i].x || (sorted_arr[j].x == arr[i].x && sorted_arr[j].y >= arr[i].y)
		})
		ans[i] = r - l
	}

	return ans
}

func gcd(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
