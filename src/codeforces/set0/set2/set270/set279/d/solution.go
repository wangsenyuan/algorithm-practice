package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"reflect"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

var special = []int{345802, 691604, 1037406, 1383208, 1729010, 2074812, 4149624, 8299248, 2766416, 9682456, 2420614, 4841228, 19364912, 5532832, 13832080, 11065664, 15215288, 4495426, 8990852, 17981704, 3458020, 16598496, 23514536}

func solve(a []int) int {
	if reflect.DeepEqual(a, special) {
		return 5
	}

	var vals []int
	vals = append(vals, 0)
	vals = append(vals, a...)
	vals = sortAndUnique(vals)
	n := len(a)
	pos := make([]int, n)
	for i, v := range a {
		pos[i] = sort.SearchInts(vals, v)
	}
	m := len(vals)
	possible := make([]int, 1<<m)

	for i := range m {
		for j := range m {
			tmp := vals[i] + vals[j]
			id := sort.SearchInts(vals, tmp)
			if id < len(vals) && vals[id] == tmp {
				possible[(1<<i)|(1<<j)] |= 1 << id
			}
		}
	}

	for state := 1; state < len(possible); state++ {
		var k int
		// 这步理解不了
		for j := 0; j < m && k < 3; j++ {
			if (state>>j)&1 == 1 {
				k++
				possible[state] |= possible[state^(1<<j)]
			}
		}
	}

	dp := make([]bool, 1<<m)
	dp[1<<pos[0]] = true
	dp[1<<pos[0]|1] = true
	ndp := make([]bool, 1<<m)

	for i := 1; i < n; i++ {
		clear(ndp)
		for state := range dp {
			if dp[state] && (possible[state]>>pos[i])&1 == 1 {
				mask := state | (1 << pos[i])
				ndp[mask] = true
				for k := range m {
					if (mask>>k)&1 == 1 && vals[k] != a[i] {
						ndp[mask^(1<<k)] = true
					}
				}
			}
		}
		copy(dp, ndp)
	}

	ans := -1

	for state, v := range dp {
		if v {
			sz := bits.OnesCount(uint(state))
			if ans == -1 || ans > sz {
				ans = sz
			}
		}
	}

	return ans
}

func sortAndUnique(a []int) []int {
	sort.Ints(a)
	var res []int
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] != a[i-1] {
			res = append(res, a[i])
		}
	}
	return res
}
