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
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	ips := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &ips[i])
	}
	return solve(k, ips)
}

func solve(k int, ips []string) string {
	n := len(ips)
	arr := make([]int, n)
	for i := range n {
		arr[i] = convert(ips[i])
	}

	play := func(mask int) int {
		res := make(map[int]bool)
		for _, ip := range arr {
			res[ip&mask] = true
		}
		return len(res)
	}

	var mask int
	for d := 31; d > 0; d-- {
		mask |= 1 << d
		cnt := play(mask)
		if cnt > k {
			break
		}
		if cnt == k {
			return format(mask)
		}
	}

	return "-1"
}

func convert(ip string) int {
	var res int
	ss := strings.Split(ip, ".")

	var shift int
	for i := 3; i >= 0; i-- {
		v, _ := strconv.Atoi(ss[i])
		res |= v << shift
		shift += 8
	}
	return res
}

func format(mask int) string {

	w0 := mask & 0xFF
	w1 := (mask >> 8) & 0xFF
	w2 := (mask >> 16) & 0xFF
	w3 := (mask >> 24) & 0xFF

	return fmt.Sprintf("%d.%d.%d.%d", w3, w2, w1, w0)
}
