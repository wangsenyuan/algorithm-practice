package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	ask := func(i int) int {
		fmt.Printf("? %d\n", i)
		return readNum(reader)
	}

	for range tc {
		n, k := readTwoNums(reader)
		res := solve(n, k, ask)
		if len(res) == 0 {
			fmt.Printf("! -1\n")
		} else {
			fmt.Printf("! %d %d\n", res[0], res[1])
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

func solve(n int, k int, ask func(int) int) []int {
	if n == 2*k {
		return []int{k, k}
	}
	arr := make([]int, n)
	// n > 2 * k
	// ask a
	for i := range k {
		arr[i] = ask(i + 1)
	}
	for i := n - k; i < n; i++ {
		arr[i] = ask(i + 1)
	}

	diffAt := n
	for i := n - k; i < n; i++ {
		if arr[i] != arr[i%k] {
			diffAt = i
			break
		}
	}
	if diffAt == n {
		return nil
	}
	var pos []int
	for i := diffAt; i >= k; i -= k {
		pos = append(pos, i)
	}

	slices.Reverse(pos)

	query := func(p int) int {
		if arr[p] == 0 {
			arr[p] = ask(p + 1)
		}
		return arr[p]
	}

	r := sort.Search(len(pos), func(i int) bool {
		return query(pos[i]) == arr[diffAt]
	})
	r = pos[r]

	l := max(k-1, r-k)
	r = min(r, n-k)

	arr[l] = arr[l%k]

	for i := l + 1; i < r+k && i < n; i++ {
		query(i)
	}

	check := func(p int) bool {
		// p是B的起点
		for j := l; j < p; j++ {
			if arr[j] != arr[j%k] {
				return false
			}
		}
		// 如果p是B的起点，
		for j := n - k; j < n; j++ {
			j1 := (j-p)%k + p
			if arr[j] != arr[j1] {
				return false
			}
		}

		return true
	}

	var ans []int

	for i := l + 1; i <= r && i <= n-k; i++ {
		if check(i) {
			ans = append(ans, i)
		}
	}

	if len(ans) != 1 {
		return nil
	}
	return []int{ans[0], n - ans[0]}
}
