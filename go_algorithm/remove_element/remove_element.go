package remove_element

import "fmt"

// 给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
// 示例1:
//     	给定 nums = [3,2,2,3], val = 3,
// 		函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
// 		你不需要考虑数组中超出新长度后面的元素。
// 示例2:
// 		给定 nums = [0,1,2,2,3,0,4,2], val = 2,
// 		函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
//		注意这五个元素可为任意顺序。
// 		你不需要考虑数组中超出新长度后面的元素。

// Methods: 1
func removeElement1(nums []int, val int) int {
	i := 0
	for index, item := range nums {
		if item != val {
			nums[i] = nums[index]
			i++
		}
	}
	fmt.Println("nums: ", nums)
	return i
}

// Methods: 2
// nums = [0,1,2,2,3,0,4,2], val = 2,
func removeElement2(nums []int, val int) int {
	i := 0
	length := len(nums)
	for {
		if i >= length {
			break
		}
		if nums[i] == val {
			nums[i] = nums[length - 1]
			length--
		} else {
			i++
		}
	}
	fmt.Println("nums: ", nums)
	return length
}

// Methods: 3
func removeElement3(nums []int, val int) int {
	i := 0
	length := len(nums)
	for {
		if i >= length {
			break
		}
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			length--
		} else {
			i++
		}
	}
	return length
}