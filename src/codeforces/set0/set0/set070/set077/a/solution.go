package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInts(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i, x := range ss {
		res[i], _ = strconv.Atoi(x)
	}
	return res
}

func drive(reader *bufio.Reader) []int {
	n := readInts(reader)[0]
	likes := make([]string, n)
	for i := range n {
		likes[i] = readString(reader)
	}
	monsters := readInts(reader)
	return solve(likes, monsters)
}

const inf = 1 << 60

func solve(likes []string, monsters []int) []int {
	// difference 只和a, b, c有关系
	slices.Sort(monsters)
	// 假设 x + y + z = 7
	best := inf
	var todo [][]int
	a, b, c := monsters[0], monsters[1], monsters[2]
	check := func(arr []int) {
		for {
			x, y, z := arr[0], arr[1], arr[2]
			u := min(a/x, b/y, c/z)
			v := max(a/x, b/y, c/z)

			if v-u < best {
				best = v - u
				todo = [][]int{arr}
			} else if v-u == best {
				todo = append(todo, arr)
			}

			if !nextPermutation(arr) {
				break
			}
		}
	}

	for x := 1; x < 7; x++ {
		for y := x; x+y < 7; y++ {
			z := 7 - (x + y)
			if y <= z {
				// [x, y, z] 去找出最好的diff
				check([]int{x, y, z})
			}
		}
	}

	heros := map[string]int{
		"Anka":        0,
		"Chapay":      1,
		"Cleo":        2,
		"Troll":       3,
		"Dracul":      4,
		"Snowy":       5,
		"Hexadecimal": 6,
	}

	g := make([]map[int]int, 7)
	for i := range 7 {
		g[i] = make(map[int]int)
	}

	for _, cur := range likes {
		ss := strings.Split(cur, " ")
		u := ss[0]
		v := ss[2]
		g[heros[u]][heros[v]] = 1
	}

	count := func(team []int) int {
		var res int
		for _, u := range team {
			for _, v := range team {
				res += g[u][v]
			}
		}
		return res
	}

	play := func(arr []int) int {
		ids := []int{0, 1, 2, 3, 4, 5, 6}
		x, y := arr[0], arr[1]
		var res int
		for {
			cnt := count(ids[:x])
			cnt += count(ids[x : x+y])
			cnt += count(ids[x+y:])

			res = max(res, cnt)

			if !nextPermutation(ids) {
				break
			}
		}
		return res
	}

	var cnt int

	for _, cur := range todo {
		cnt = max(cnt, play(cur))
	}

	return []int{best, cnt}
}

func nextPermutation(arr []int) bool {
	// Find longest decreasing suffix
	i := len(arr) - 2
	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}

	if i < 0 {
		return false // No next permutation
	}

	// Find successor to pivot in suffix
	j := len(arr) - 1
	for arr[j] <= arr[i] {
		j--
	}

	// Swap pivot with successor
	arr[i], arr[j] = arr[j], arr[i]

	// Reverse the suffix
	for k := 1; i+k < len(arr)-k; k++ {
		arr[i+k], arr[len(arr)-k] = arr[len(arr)-k], arr[i+k]
	}

	return true
}
