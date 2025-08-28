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
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	s := make([]string, n)
	for i := range n {
		s[i] = readString(reader)
	}
	m := readNum(reader)
	ops := make([][]int, m)
	for i := range m {
		ops[i] = readNNums(reader, 2)
	}
	return solve(s, ops)
}

const H = 9

func solve(s []string, ops [][]int) []int {
	n := len(s)

	// fp[i]表示s[i]的最大的k
	fp := make([]int, n)

	pref := make([]string, n)
	suf := make([]string, n)
	width := make([]int, n)

	getFlag := func(s string, k int) *BitSet {
		flag := NewBitSet(1 << k)
		if len(s) < 1<<k {
			return flag
		}
		var num int
		for i := 0; i < k; i++ {
			num = num*2 + int(s[i]-'0')
		}
		flag.Set(num)

		for i := k; i < len(s); i++ {
			if num&(1<<(k-1)) > 0 {
				// 去掉最高位
				num ^= 1 << (k - 1)
			}
			num <<= 1
			num += int(s[i] - '0')
			flag.Set(num)
		}
		return flag
	}

	calc := func(s string) int {
		for k := 1; k <= H; k++ {
			flag := getFlag(s, k)
			if flag.Count() != 1<<k {
				return k - 1
			}
		}
		return H
	}
	K := 1 << H

	getPrefix := func(s string) string {
		if len(s) <= K {
			return s
		}
		return s[:K]
	}

	getSuffix := func(s string) string {
		if len(s) <= K {
			return s
		}
		return s[len(s)-K:]
	}

	for i := 0; i < n; i++ {
		fp[i] = calc(s[i])
		pref[i] = getPrefix(s[i])
		suf[i] = getSuffix(s[i])
		width[i] = len(s[i])
	}

	ans := make([]int, len(ops))

	for i, cur := range ops {
		a, b := cur[0]-1, cur[1]-1
		k := max(fp[a], fp[b])
		w := suf[a] + pref[b]
		// 在k的基础上增加
		for k+1 <= H {
			flag := getFlag(w, k+1)
			if flag.Count() < 1<<(k+1) {
				break
			}
			k++
		}
		fp = append(fp, k)
		ans[i] = k
		if len(pref[a]) == width[a] {
			pref = append(pref, getPrefix(w))
		} else {
			pref = append(pref, pref[a])
		}
		if len(suf[b]) == width[b] {
			suf = append(suf, getSuffix(w))
		} else {
			suf = append(suf, suf[b])
		}
		width = append(width, min(2*K, width[a]+width[b]))
	}

	return ans
}

type BitSet struct {
	sz  int
	arr []uint64
}

func NewBitSet(n int) *BitSet {
	sz := (n + 63) / 64
	arr := make([]uint64, sz)
	return &BitSet{sz, arr}
}

func (bs *BitSet) Set(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	bs.arr[i] |= 1 << uint64(j)
}

func (bs *BitSet) IsSet(p int) bool {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	return (bs.arr[i]>>uint64(j))&1 == 1
}

func (bs *BitSet) Count() int {
	var res int
	for i := 0; i < bs.sz; i++ {
		res += countDigit(bs.arr[i])
	}
	return res
}

func countDigit(num uint64) int {
	var res int
	if (num>>63)&1 == 1 {
		res++
	}
	tmp := int64(num & ((1 << 63) - 1))

	for tmp > 0 {
		res++
		tmp -= tmp & -tmp
	}
	return res
}

func (bs *BitSet) LeftShift(p int) {
	i, j := p/64, p%64
	if j == 0 {
		for u := 0; u+i < bs.sz; u++ {
			bs.arr[u] = bs.arr[u+i]
		}
	} else {
		for u := 0; u+i < bs.sz; u++ {
			v := u + i
			bs.arr[u] = bs.arr[v] << uint64(j)
			if v+1 < bs.sz {
				bs.arr[u] |= bs.arr[v+1] >> uint64(64-j)
			}
		}
	}
	for u := bs.sz - i; u < bs.sz; u++ {
		bs.arr[u] = 0
	}
}

func (bs *BitSet) RightShift(p int) {
	i, j := p/64, p%64
	if j == 0 {
		for u := bs.sz - 1; u-i >= 0; u-- {
			bs.arr[u] = bs.arr[u-i]
		}
	} else {
		for u := bs.sz - 1; u-i >= 0; u-- {
			v := u - i
			bs.arr[u] = bs.arr[v] >> uint64(j)
			if v-1 >= 0 {
				bs.arr[u] |= bs.arr[v-1] << uint64(64-j)
			}
		}
	}
	for u := i - 1; u >= 0; u-- {
		bs.arr[u] = 0
	}
}

func (bs *BitSet) Copy() *BitSet {
	res := NewBitSet(len(bs.arr) * 64)
	copy(res.arr, bs.arr)
	return res
}

func (this *BitSet) Union(that *BitSet) {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] |= that.arr[i]
	}
}
