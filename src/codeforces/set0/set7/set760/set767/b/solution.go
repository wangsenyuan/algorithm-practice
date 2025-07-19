package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, res := process(reader)
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

func process(reader *bufio.Reader) (int, int, int, []int, int) {
	ts, tf, t := readThreeNums(reader)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(ts, tf, t, a)
	return ts, tf, t, a, res
}

const inf = 1 << 60

func solve(ts int, tf int, t int, a []int) int {
	if len(a) == 0 || a[0] > ts {
		return ts
	}
	// a is already sorted
	// a is positive
	ans := []int{ts - a[0] + 1, a[0] - 1}

	timeline := ts

	n := len(a)

	for i := 0; i < n; i++ {
		if a[i]+t > tf {
			// 在a[i]时刻到达，没有足够的时间工作了
			timeline = tf
			break
		}
		// 如果在a[i]的前面到达呢？
		if i > 0 && a[i]-1 > a[i-1] {
			wait := max(0, timeline-(a[i]-1))
			if wait < ans[0] {
				ans[0] = wait
				ans[1] = a[i] - 1
			}
		}

		timeline = max(timeline, a[i]) + t
		if i == n-1 || a[i] > a[i+1] {
			// 可以在这个时刻到达
			if timeline-a[i] < ans[0] {
				ans[0] = timeline - a[i]
				ans[1] = a[i]
			}
		}
	}

	if timeline+t <= tf {
		// 最后一个，不用等
		ans[0] = 0
		ans[1] = timeline
	}

	return ans[1]
}
