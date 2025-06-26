package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		_, res := process(reader)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
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

func process(reader *bufio.Reader) (a []int, res []int) {
	n := readNum(reader)
	a = readNNums(reader, n)
	res = solve(n, a)
	return
}

func solve(n int, a []int) []int {
	// 如果i被保留，表示 b[i-1] > b[i] < b[i+1]
	// 如果i被删除，那么要么 b[i-1] < b[i], 或者 b[i+1] < b[i]
	// 如果是连续的一段被删除的话，必然是一段递增或者递减，没有其他的可能性

	arr := make([]int, n)
	for i := range n {
		arr[i] = i
		if a[i] == -1 {
			a[i] = n
		}
	}
	flag := make([]bool, n)
	id := 1
	l, r := 1, n
	ans := make([]int, n)

	for len(arr) > 1 {
		for j, i := range arr {
			if a[i] == id {
				// 这些被删除了
				flag[j] = true
			}
		}

		if id%2 == 1 {
			for j := 0; j < len(arr); {
				if !flag[j] {
					// 这个保留的
					j++
					continue
				}
				k := j
				for j < len(arr) && flag[j] {
					j++
				}

				if k == 0 {
					for k < j {
						ans[arr[k]] = r
						r--
						k++
					}
				} else {
					w := r - (j - k) + 1
					r -= j - k
					for k < j {
						ans[arr[k]] = w
						w++
						k++
					}
				}
			}
		} else {
			for j := 0; j < len(arr); {
				if !flag[j] {
					j++
					continue
				}
				k := j
				for j < len(arr) && flag[j] {
					j++
				}
				if k == 0 {
					for k < j {
						ans[arr[k]] = l
						l++
						k++
					}
				} else {
					w := l + j - k - 1
					l += j - k
					for k < j {
						ans[arr[k]] = w
						w--
						k++
					}
				}
			}
		}
		var next []int
		for j := range len(arr) {
			if !flag[j] {
				next = append(next, arr[j])
			} else {
				flag[j] = false
			}
		}
		arr = next
		id++
	}

	ans[arr[0]] = l

	return ans
}
