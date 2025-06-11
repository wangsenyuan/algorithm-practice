package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)
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

const mod = 1000000007

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n int) int {
	if n == 1 {
		return 1
	}

	sum := []int{2, 1}

	for i := 3; i <= n; i++ {
		cur := add(sum[(i-1)&1], 1)
		sum[i&1] = add(sum[i&1], cur)
	}

	return add(sum[0], sum[1])
}
