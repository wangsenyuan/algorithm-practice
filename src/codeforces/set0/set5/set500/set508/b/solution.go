package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	num := readString(reader)
	res := solve(num)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(num string) string {
	// 比如交换一个偶数到最后的位置
	n := len(num)
	// 要么把第一个小于最后一个数的偶数交换过去
	y := int(num[n-1] - '0')
	buf := []byte(num)
	for i := 0; i < n-1; i++ {
		x := int(num[i] - '0')
		if x&1 == 0 && x < y {
			buf[i], buf[n-1] = buf[n-1], buf[i]
			return string(buf)
		}
	}
	// 没有比y小的偶数， 就是最后一个偶数
	for i := n - 1; i >= 0; i-- {
		x := int(num[i] - '0')
		if x&1 == 0 {
			buf[i], buf[n-1] = buf[n-1], buf[i]
			return string(buf)
		}
	}
	return "-1"
}
