package main

import (
	"./fizz_buzz"
	"./merge_sorted_array"
	"./remove_duplicates"
	"./single_number"
	"./three_sums"
	"fmt"
)

func main() {
	testData1 := []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	//testData2 := []int{-1, 0, 1, 1, 2, -1, -4}
	//fmt.Println("testData: ", testData)
	fmt.Println(three_sums.ThreeSum(testData1))
	testSingleNumber()
	testRemoveDuplicates()
	testFizzBuzz()
	testMergeSortedArray()
}

func testSingleNumber() {
	testData := []int{1, 2, 3, 2, 3, 1, 5}
	res := single_number.SingleNumber(testData)
	fmt.Println(res)
}

func testRemoveDuplicates() {
	nums := []int{0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := remove_duplicates.RemoveDuplicates(nums)
	fmt.Println("Duplicates: ", res)
}

func testFizzBuzz() {

	res := fizz_buzz.FizzBuzz(15)
	fmt.Println("FizzBuzz: ", res)
}

func testMergeSortedArray() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge_sorted_array.Merge(nums1, 3, nums2, 3)
	fmt.Println("mergeSortedArray: ", nums1)
}
