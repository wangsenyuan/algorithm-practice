package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(strconv.Itoa(x))
		buf.WriteByte('\n')
	}
	writer.Write(buf.Bytes())
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	pos := make([]int, n+1)
	for i := range n {
		var v int
		fmt.Fscan(reader, &v)
		pos[v] = i
	}

	for i, v := range a {
		a[i] = pos[v]
	}

	wm := newWaveletMatrix(a, n)

	var m int
	fmt.Fscan(reader, &m)
	ans := make([]int, m)
	var x int
	f := func(z int) int {
		return ((z - 1 + x) % n) + 1
	}

	for i := range m {
		var b1, b2, c1, c2 int
		fmt.Fscan(reader, &b1, &b2, &c1, &c2)

		l1, r1 := f(b1), f(b2)
		l1, r1 = min(l1, r1), max(l1, r1)
		l2, r2 := f(c1), f(c2)
		l2, r2 = min(l2, r2), max(l2, r2)

		ans[i] = wm.rangeFreq(l1-1, r1, l2-1, r2)
		x = ans[i] + 1
	}

	return ans
}

type bitVector struct {
	bits []uint64
	pref []uint32
}

func newBitVector(n int) bitVector {
	sz := (n + 63) >> 6
	return bitVector{
		bits: make([]uint64, sz),
		pref: make([]uint32, sz+1),
	}
}

func (bv *bitVector) set(i int) {
	bv.bits[i>>6] |= 1 << uint(i&63)
}

func (bv *bitVector) build() {
	for i, v := range bv.bits {
		bv.pref[i+1] = bv.pref[i] + uint32(bits.OnesCount64(v))
	}
}

func (bv *bitVector) rank1(r int) int {
	word := r >> 6
	if word >= len(bv.bits) {
		return int(bv.pref[len(bv.bits)])
	}

	res := int(bv.pref[word])
	if bit := r & 63; bit > 0 {
		res += bits.OnesCount64(bv.bits[word] & ((uint64(1) << uint(bit)) - 1))
	}
	return res
}

func (bv *bitVector) rank0(r int) int {
	return r - bv.rank1(r)
}

type waveletMatrix struct {
	lg   int
	mid  []int
	data []bitVector
}

func newWaveletMatrix(a []int, upper int) *waveletMatrix {
	lg := bits.Len(uint(upper))
	if lg == 0 {
		lg = 1
	}

	cur := a
	next := make([]int, len(a))
	wm := &waveletMatrix{
		lg:   lg,
		mid:  make([]int, lg),
		data: make([]bitVector, lg),
	}

	for level := 0; level < lg; level++ {
		shift := lg - 1 - level
		bv := newBitVector(len(a))
		var zeroCount int
		for i, v := range cur {
			if (v>>shift)&1 == 1 {
				bv.set(i)
			} else {
				zeroCount++
			}
		}
		bv.build()
		wm.data[level] = bv
		wm.mid[level] = zeroCount

		p0, p1 := 0, zeroCount
		for _, v := range cur {
			if (v>>shift)&1 == 0 {
				next[p0] = v
				p0++
			} else {
				next[p1] = v
				p1++
			}
		}
		cur, next = next, cur
	}

	return wm
}

func (wm *waveletMatrix) countLess(l int, r int, upper int) int {
	if l >= r || upper <= 0 {
		return 0
	}
	if upper >= 1<<wm.lg {
		return r - l
	}

	var res int
	for level := 0; level < wm.lg; level++ {
		bv := &wm.data[level]
		l0, r0 := bv.rank0(l), bv.rank0(r)
		if (upper>>(wm.lg-1-level))&1 == 1 {
			res += r0 - l0
			l = wm.mid[level] + (l - l0)
			r = wm.mid[level] + (r - r0)
		} else {
			l, r = l0, r0
		}
	}

	return res
}

func (wm *waveletMatrix) rangeFreq(l int, r int, lower int, upper int) int {
	if l >= r || lower >= upper {
		return 0
	}
	return wm.countLess(l, r, upper) - wm.countLess(l, r, lower)
}
