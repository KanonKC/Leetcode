package main

import "fmt"

func removeElement(nums []int, val int) int {
	
	writeIndex := 0

	for readIndex, _ := range nums {
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	return writeIndex
}

func main() {
	nums := []int{3,2,2,3}
	val := 2
	result := removeElement(nums, val)
	fmt.Println(result)
	fmt.Println(nums)
}