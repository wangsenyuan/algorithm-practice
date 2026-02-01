package p3827

func countMonobit(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	res := 1
	for i := 1; 1<<i-1 <= n; i++ {
		res++
	}
	return res
}
