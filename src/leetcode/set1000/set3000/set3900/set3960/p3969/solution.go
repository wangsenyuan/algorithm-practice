package p3969

func countValidSubarrays(nums []int, x int) int {

	var res int

	for i := range nums {
		var sum int
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum%10 == x && getHigh(sum) == x {
				res++
			}
		}
	}

	return res
}

func getHigh(num int) int {
	for num >= 10 {
		num /= 10
	}
	return num
}
