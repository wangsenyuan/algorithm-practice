package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	beauties := readNNums(reader, n)
	intellects := readNNums(reader, n)
	richnesses := readNNums(reader, n)
	return solve(beauties, intellects, richnesses)
}

type Person struct {
	beauty, intellect, richness int
}

func solve(beauties []int, intellect []int, richness []int) int {
	var intellects []int
	n := len(beauties)
	arr := make([]Person, n)
	for i := range n {
		intellects = append(intellects, intellect[i])
		arr[i] = Person{beauties[i], intellect[i], richness[i]}
	}
	intellects = sortAndUnique(intellects)
	m := len(intellects)
	slices.SortFunc(arr, func(a, b Person) int {
		return b.beauty - a.beauty
	})
	tr := make(SegmentTree, 2*m)

	var ans int

	for i := 0; i < n; {
		j := i
		for i < n && arr[i].beauty == arr[j].beauty {
			i1 := sort.SearchInts(intellects, arr[i].intellect)
			x := tr.Query(i1+1, m)
			if x > arr[i].richness {
				ans++
			}
			i++
		}
		for j < i {
			cur := arr[j]
			i1 := sort.SearchInts(intellects, cur.intellect)
			tr.Update(i1, cur.richness)
			j++
		}
	}
	return ans
}

func sortAndUnique(arr []int) []int {
	sort.Ints(arr)
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] != arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
}

type SegmentTree []int

func (t SegmentTree) Update(p int, v int) {
	n := len(t) / 2
	p += n
	t[p] = max(t[p], v)
	for p > 1 {
		t[p>>1] = max(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegmentTree) Query(l int, r int) int {
	n := len(t) / 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
