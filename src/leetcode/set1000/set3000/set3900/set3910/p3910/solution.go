package p3910

func evenSumSubgraphs(nums []int, edges [][]int) int {
	n := len(nums)
	set := NewDSU(n)

	var res int

	for mask := 1; mask < (1 << n); mask++ {
		var sum int
		first := -1
		for i := range n {
			if (mask>>i)&1 == 1 {
				sum += nums[i]
				first = i
			}
		}
		if sum%2 == 1 {
			continue
		}
		for _, e := range edges {
			u, v := e[0], e[1]
			if (mask>>u)&1 == 1 && (mask>>v)&1 == 1 {
				set.Union(u, v)
			}
		}
		ok := true
		for i := range n {
			if (mask>>i)&1 == 1 {
				if set.Find(i) != set.Find(first) {
					ok = false
					break
				}
			}
		}
		if ok {
			res++
		}
		set.Reset()
	}
	return res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

func (this *DSU) Reset() {
	for i := range this.arr {
		this.arr[i] = i
		this.cnt[i] = 1
	}
}
