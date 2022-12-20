package easy

func RemoveElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		// 循环遍历nums，满足条件的，重新刷入nums
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
