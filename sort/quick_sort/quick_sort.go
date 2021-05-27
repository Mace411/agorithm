package main

import "fmt"

func main() {
	var nums = []int{8, 7, 6, 5, 4, 3, 2, 1 /*3, 2, 4, 7, 8, 0, 1, 4, 2, 6*/ /*1, 2, 3, 4 ,5, 6, 7, 8/*, 99,8,7,6,5,4,3,2,1*/} //8,7,6,5,4,3,2,1
	//for i := 0; i < len(nums)/2; i++ {
	//	nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	//}
	//fmt.Printf("%v\n", nums)
	quickSort(nums, 0, len(nums)-1)
	fmt.Print(nums)
}

//func partition(nums []int, low, high int) int {
//	var p = nums[high]
//	var i = low - 1
//	for j := low; j < high; j++ {
//		if nums[j] < p {
//			i++
//			nums[i], nums[j] = nums[j], nums[i]
//		}
//	}
//	nums[i+1], nums[high] = nums[high], nums[i+1]
//	return i + 1
//}

//有序的情况就变成了O(n^2)，逆序的情况
func partition(nums []int, low, high int) int {
	var p = nums[low]
	for low < high {
		for nums[high] >= p && low < high {
			high--
		}
		nums[low], nums[high] = nums[high], nums[low]
		for nums[low] <= p && low < high {
			low++
		}
		nums[low], nums[high] = nums[high], nums[low]
	}
	return low
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
