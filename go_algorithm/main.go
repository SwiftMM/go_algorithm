package main

import (
	"./three_sums"
	"fmt"
)

func main() {
	testData1 := []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	//testData2 := []int{-1, 0, 1, 1, 2, -1, -4}
	//fmt.Println("testData: ", testData)
	fmt.Println(three_sums.ThreeSum(testData1))

}
