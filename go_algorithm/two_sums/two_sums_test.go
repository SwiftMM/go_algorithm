package two_sums

import "testing"

func TestTwoSums(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	s := twoSum(nums, target)
	if !equal([]int{0,1}, s) {
		t.Error("test failed")
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}