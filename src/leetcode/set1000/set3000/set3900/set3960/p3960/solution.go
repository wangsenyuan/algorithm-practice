package p3960

func getLength(nums []int) int {

	var best int

	for i := range nums {
		f1 := make(map[int]int)
		f2 := make(map[int]int)
		for j := i; j >= 0; j-- {
			v := nums[j]
			if f1[v] > 0 {
				f2[f1[v]]--
				if f2[f1[v]] == 0 {
					delete(f2, f1[v])
				}
			}
			f1[v]++
			f2[f1[v]]++
			if len(f1) == 1 {
				best = max(best, i-j+1)
			}
			if len(f2) == 2 {
				var arr []int
				for k := range f2 {
					arr = append(arr, k)
				}
				if arr[0] == arr[1]*2 || arr[0]*2 == arr[1] {
					best = max(best, i-j+1)
				}
			}
		}
	}

	return best
}
