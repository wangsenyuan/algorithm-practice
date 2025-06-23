package p3587

func minSwaps(nums []int) int {
	var odd []int
	var even []int

	for i, v := range nums {
		if v%2 == 1 {
			odd = append(odd, i)
		} else {
			even = append(even, i)
		}
	}

	if abs(len(odd)-len(even)) > 1 {
		return -1
	}
	n := len(nums)
	arr := make([]int, n)
	check := func(a, b []int) int {
		for i := 0; i < n; i += 2 {
			arr[a[0]] = i
			a = a[1:]
			if len(b) > 0 {
				arr[b[0]] = i + 1
				b = b[1:]
			}
		}
		return countInversions(arr)
	}

	ans := 1 << 60
	if len(odd) >= len(even) {
		ans = check(odd, even)
	}
	if len(even) >= len(odd) {
		ans = min(ans, check(even, odd))
	}
	return ans
}

func countInversions(arr []int) int {
	n := len(arr)
	set := make(BIT, n+3)
	var res int
	for i := 0; i < n; i++ {
		res += i - set.Get(arr[i])
		set.Update(arr[i], 1)
	}
	return res
}

type BIT []int

func (set BIT) Update(p int, v int) {
	p++
	for p < len(set) {
		set[p] += v
		p += p & -p
	}
}

func (set BIT) Get(p int) int {
	p++
	var res int
	for p > 0 {
		res += set[p]
		p -= p & -p
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
