package easy

func RemoveDuplicates(nums []int) int {
	i := 0
	// 双指针，快慢赋值
	for j := 1; j < len(nums); j++ {
		if nums[i] == nums[j] {
			continue
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
