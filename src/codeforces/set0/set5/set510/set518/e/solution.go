package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	if len(res) == 0 {
		fmt.Println("Incorrect sequence")
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func process(reader *bufio.Reader) (int, []int) {
	_, k := readTwoNums(reader)
	s := readString(reader)
	return k, solve(k, s)
}

const inf = 1 << 60

func solve(k int, s string) []int {
	nums := strings.Split(s, " ")
	n := len(nums)
	a := make([]int, n)
	mark := make([]bool, n)

	for i := range n {
		if nums[i] == "?" {
			mark[i] = true
		} else {
			a[i], _ = strconv.Atoi(nums[i])
		}
	}

	arr := make([]int, (n+k-1)/k)

	assign := func(pos int, hi int) bool {
		lo := -inf
		if arr[0]-k >= 0 {
			lo = a[arr[0]-k]
		}
		hi--
		lo++
		if hi-lo+1 < pos {
			// 这个空间不够放一个递增序列
			return false
		}
		// 否则肯定有答案
		if hi <= 0 {
			for i := pos - 1; i >= 0; i-- {
				a[arr[i]] = hi
				hi--
			}
		} else if lo >= 0 {
			for i := 0; i < pos; i++ {
				a[arr[i]] = lo
				lo++
			}
		} else {
			// lo < 0 < hi,
			// 假设最优的区间是 x.0..y
			// 那么 x >= lo, y <= hi, 且x和y
			// y - x + 1 = pos
			// 假设 -x = y
			y := (pos - 1) / 2
			x := y + 1 - pos
			if hi < y {
				x -= y - hi
				y = hi
			} else if x < lo {
				y += lo - x
				x = lo
			}
			for i := 0; i < pos; i++ {
				a[arr[i]] = x + i
			}
		}

		return true
	}

	for i := 0; i < k; i++ {
		var pos int
		for j := i; j < n; j += k {
			if mark[j] {
				arr[pos] = j
				pos++
				continue
			}
			// !mark[j]
			if pos > 0 && !assign(pos, a[j]) {
				return nil
			}
			pos = 0
		}
		if pos > 0 && !assign(pos, inf) {
			return nil
		}
	}

	for i := k; i < n; i++ {
		if a[i-k] >= a[i] {
			return nil
		}
	}

	return a
}
