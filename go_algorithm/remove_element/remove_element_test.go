package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	target := 3
	s := RemoveElement1(nums, target)
	if s != 2 {
		t.Error("test failed")
	}
}
