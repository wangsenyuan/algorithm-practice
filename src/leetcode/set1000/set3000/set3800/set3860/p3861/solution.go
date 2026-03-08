package p3861

func minimumIndex(capacity []int, itemSize int) int {
	n := len(capacity)
	ans := -1
	for i := n - 1; i >= 0; i-- {
		if capacity[i] < itemSize {
			continue
		}
		if ans < 0 || capacity[i] <= capacity[ans] {
			ans = i
		}
	}
	return ans
}
