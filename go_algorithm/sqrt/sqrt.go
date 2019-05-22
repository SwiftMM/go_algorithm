package sqrt

import "fmt"

// 二分查找法
func MySqrt(x int) int {

	if x <= 0 {
		return 0
	}
	i := 1
	j := x/2 + 1
	for i <= j {
		mid := (i + j) / 2
		res := mid * mid
		if res == x {
			return mid
		} else if res < x {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return j
}

// 牛顿法
func MySqrt2(x int) int {
	if x <= 0 {
		return 0
	}

	var last = 0.0
	var res = 1.0
	fmt.Println("MySqrt2")
	for res != last {
		last = res
		res = (res + float64(x)/res) / 2.0
		fmt.Println(last, res)
	}
	return int(res)
}
