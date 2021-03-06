package reverseInteger

import (
	"testing"
)

var testData = map[int]int{
	123:        321,
	-123:       -321,
	120:        21,
	1534236469: 0,
}

func TestReverseInteger(t *testing.T) {
	for k, v := range testData {
		testReverseInteger(k, v, t)
	}
}

func testReverseInteger(input, output int, t *testing.T) {
	s := Reverse(input)
	if s != output {
		t.Errorf("test failed with input: %v, ouput: %v", input, output)
	}
}
