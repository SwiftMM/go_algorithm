package remove_duplicates

func RemoveDuplicates(nums []int) int {
	i := 0
	j := 1
	l := len(nums)
	if l <= 1 {
		return l
	}
	for {
		if j >= l {
			break
		}
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}
	return i + 1
}
