package main

import (
	"fmt"
)

// slice- dynamic array
// most construct in go, useful methods
func main() {
	// uninitialized slice is nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 1, 5)
	// nums := []int{}
	fmt.Println(nums)
	// capacity-> maximum number of elements can fit
	// fmt.Println(cap(nums))
	// nums[0] = 4
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	//copy function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 4)
	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//slice opearator

	// var nums = []int{1, 2, 3}

	// fmt.Println(nums[1:3])
	// fmt.Println(nums[:3])
	// fmt.Println(nums[1:])

	// slice
	// var nums = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums, nums2))

	// 2d slices
	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)

}
