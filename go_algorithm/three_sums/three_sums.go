package three_sums

import (
	"sort"
)

// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 示例:
//		例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
// 		满足要求的三元组集合为：
//		[
//  		[-1, 0, 1],
// 			[-1, -1, 2]
//		]

func ThreeSum(nums []int) [][]int {
	length := len(nums)
	res := make([][]int, 0)
	// 特殊情况,数组长度小于3,返回空数组
	if length < 3 {
		return res
	}
	sort.Ints(nums)
	for index := range nums {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		i, j := index+1, length-1
		for i < j {
			sum := nums[index] + nums[i] + nums[j]
			if sum == 0 {
				array := []int{nums[index], nums[i], nums[j]}
				res = append(res, array)
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else if sum > 0 {
				j--
			} else {
				i++
			}
		}
	}
	return res
}
