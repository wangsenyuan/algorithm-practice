package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ok, res, _, _ := process(reader)
	if !ok {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, cur := range res {
		buf.WriteString(cur)
		buf.WriteByte('\n')
	}
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

func process(reader *bufio.Reader) (bool, []string, []int, []int) {
	n := readNum(reader)
	a := readNNums(reader, n)
	k := readNum(reader)
	b := readNNums(reader, k)
	ok, res := solve(a, b)
	return ok, res, a, b
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(a []int, b []int) (bool, []string) {
	var res []string

	var offset int
	var j int
	for i := 0; i < len(b); i++ {
		j1 := j
		var sum, cnt int
		mx, mn := j, j
		for j < len(a) && sum < b[i] {
			sum += a[j]
			if a[j] > a[mx] {
				mx = j
			}
			if a[j] < a[mn] {
				mn = j
			}
			j++
			cnt++
		}
		if sum != b[i] || cnt > 1 && a[mn] == a[mx] {
			return false, nil
		}

		if cnt > 1 {

			if j1 == mx {
				for a[mx] == a[mx+1] {
					mx++
				}
			}

			if mx+1 == j || a[mx] == a[mx+1] {
				// 只能处理左边的
				for k := mx - 1; k >= j1; k-- {
					res = append(res, fmt.Sprintf("%d L", k+2-offset))
				}
				offset += mx - j1
				for k := mx + 1; k < j; k++ {
					res = append(res, fmt.Sprintf("%d R", mx+1-offset))
				}
				offset += j - mx - 1
			} else {
				for k := mx + 1; k < j; k++ {
					res = append(res, fmt.Sprintf("%d R", mx+1-offset))
				}
				for k := mx - 1; k >= j1; k-- {
					res = append(res, fmt.Sprintf("%d L", k+2-offset))
				}
				offset += cnt - 1
			}
		}
	}
	if j < len(a) {
		return false, nil
	}
	return true, res
}
