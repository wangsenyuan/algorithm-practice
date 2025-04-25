package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	fmt.Println(solve(s))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')

	return strings.TrimSpace(s)
}

func solve(s string) int {
	n := len(s)

	fp := make([]int, n+1)
	fp[n] = 0
	p := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		x := s[i:]
		kmp(x, p)
		fp[i] = min(len(x)+1, fp[i+1]+2)
		for j := 1; j < len(x); j++ {
			// k是周期
			k := j + 1 - p[j]
			if p[j] > 0 && (j+1)%k == 0 {
				fp[i] = min(fp[i], k+digitLength((j+1)/k)+fp[i+j+1])
			} else {
				fp[i] = min(fp[i], j+2+fp[i+j+1])
			}
		}
	}

	return fp[0]
}

func digitLength(x int) int {
	var res int
	for x >= 10 {
		res++
		x /= 10
	}
	return res + 1
}

func kmp(s string, p []int) {
	clear(p)
	for i := 1; i < len(s); i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
}

const M = 1000000007

const B1 = 29
const B2 = 31

type Key struct {
	first  int
	second int
}

func (this Key) Add(v int) Key {
	first := (this.first*B1%M + v) % M
	second := (this.second*B2%M + v) % M
	return Key{first, second}
}

func (this Key) Mul(a, b int) Key {
	first := this.first * a % M
	second := this.second * b % M
	return Key{first, second}
}

func (this Key) Mul2(that Key) Key {
	first := this.first * that.first % M
	second := this.second * that.second % M
	return Key{first, second}
}

func (this Key) Sub(that Key) Key {
	first := (this.first + M - that.first) % M
	second := (this.second + M - that.second) % M
	return Key{first, second}
}
