package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func process(reader *bufio.Reader) string {
	n, m := readTwoNums(reader)
	read := func() (bool, string) {
		s, err := reader.ReadString('\n')
		return err == nil, strings.TrimSpace(s)
	}
	return solve(read, n, m)
}

const format = "2012-MM-DD HH:MM:SS"

type data struct {
	ts  string
	val int
	cnt int
}

func parseTimestamp(s string) data {
	// format is 2012-MM-DD HH:MM:SS:MESSAGE
	s = s[:len(format)]
	// MM-DD HH:MM:SS
	var month, day, hour, minute, second int
	fmt.Sscanf(s, "2012-%d-%d %d:%d:%d", &month, &day, &hour, &minute, &second)
	month--
	day--
	val := month*31*24*60*60 + day*24*60*60 + hour*60*60 + minute*60 + second
	return data{ts: s, val: val, cnt: 1}
}

func solve(read func() (bool, string), n int, m int) string {
	var que []data
	var sum int
	var tail int
	for {
		ok, s := read()
		if !ok {
			break
		}
		cur := parseTimestamp(s)
		if tail < len(que) && que[tail].val == cur.val {
			que[tail].cnt++
		} else {
			que = append(que, cur)
		}
		sum++
		for tail < len(que) && que[tail].val <= cur.val-n {
			sum -= que[tail].cnt
			tail++
		}
		if sum >= m {
			return cur.ts
		}
	}
	return "-1"
}
