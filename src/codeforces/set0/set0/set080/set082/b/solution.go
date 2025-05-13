package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d", len(cur)))
		for _, v := range cur {
			buf.WriteString(fmt.Sprintf(" %d", v))
		}
		buf.WriteString("\n")
	}
	fmt.Print(buf.String())
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

func readVarNums(reader *bufio.Reader) []int {
	s, _ := reader.ReadBytes('\n')
	var m int
	pos := readInt(s, 0, &m) + 1
	res := make([]int, m)
	for j := range m {
		pos = readInt(s, pos, &res[j]) + 1
	}
	return res
}
func process(reader *bufio.Reader) [][]int {
	n := readNum(reader)
	result := make([][]int, n*(n-1)/2)
	for i := range result {
		result[i] = readVarNums(reader)
	}
	return solve(n, result)
}

type data [4]uint64

func parseData(cur []int) data {
	var res data
	for _, v := range cur {
		i, j := v/64, v%64
		res[i] |= 1 << j
	}
	return res
}

func convert(a data) []int {
	var res []int
	for i := range a {
		for j := 0; j < 64; j++ {
			if a[i]&(1<<j) != 0 {
				res = append(res, i*64+j)
			}
		}
	}
	return res
}

func contains(a, b data) bool {
	for i := range a {
		if a[i]&b[i] != b[i] {
			return false
		}
	}
	return true
}

func xor(a, b data) data {
	for i := range a {
		a[i] ^= b[i]
	}
	return a
}

var zero data

func solve(n int, result [][]int) [][]int {
	if n == 2 {
		a := result[0][0]
		return [][]int{
			{a},
			result[0][1:],
		}
	}
	m := n * (n - 1) / 2
	arr := make([]data, m)
	var nums []int

	for i, cur := range result {
		arr[i] = parseData(cur)
		nums = append(nums, cur...)
	}
	nums = sortAndUnique(nums)
	// x <= 200
	freq2 := make([]int, 201)
	marked := make([]bool, 201)
	var res [][]int
	for _, x := range nums {
		if marked[x] {
			continue
		}
		// 找到所有和x在一起的，且freq相同的值
		// 先找出和x在一起的
		clear(freq2)
		i, j := x/64, x%64

		for _, cur := range arr {
			if cur[i]&(1<<j) != 0 {
				for _, y := range nums {
					l, r := y/64, y%64
					if cur[l]&(1<<r) != 0 {
						freq2[y]++
					}
				}
			}
		}

		var tmp []int
		for _, y := range nums {
			if freq2[y] == freq2[x] {
				tmp = append(tmp, y)
				marked[y] = true
			}
		}
		if len(res) == 0 {
			// 只有第一个需要添加进去。其他的貌似都在它被计算出来前，就已经加入了
			res = append(res, tmp)
		}
		var p int

		// 一次可以减少n-1个
		dt := parseData(tmp)

		for _, cur := range arr {
			if !contains(cur, dt) {
				arr[p] = cur
				p++
				continue
			}
			cur = xor(cur, dt)
			tmp = convert(cur)
			res = append(res, tmp)
		}
		arr = arr[:p]
		if len(res) == n {
			break
		}
	}

	return res
}

func sortAndUnique(nums []int) []int {
	sort.Ints(nums)
	var p int
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != nums[i-1] {
			nums[p] = nums[i-1]
			p++
		}
	}
	return nums[:p]
}
