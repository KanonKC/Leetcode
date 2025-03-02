package main

import "fmt"

func removeDuplicates(nums []int) int {
    prev := -999
	writeIndex := 0

	for _, num := range nums {
		if num != prev {
			nums[writeIndex] = num
			prev = num
			writeIndex++
		}
	}

	return writeIndex
}

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	k := removeDuplicates(nums)
	fmt.Println(k)
	fmt.Println(nums)
}