package main

import "fmt"

func main() {
	var nums = []int{3, 2, 4, 7, 5, 8, 0, 1, 4, 2, 6, 9 /*1, 2, 3, 4 ,5, 6, 7, 8, 99,8,7,6,5,4,3,2,1*/}
	quickSort(nums, 0, len(nums)-1)
	fmt.Print(nums)
}

//递归实现
func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(nums, low, high)
	quickSort(nums, low, p-1)
	quickSort(nums, p+1, high)
}

func partition(nums []int, low, high int) int {
	return 0
}
