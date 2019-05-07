package single_number

func SingleNumber(nums []int) int {
	a := 0
	for _, e := range nums {
		a = a ^ e
	}
	return a
}
