package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 2, 1}
	var answer bool
	j := len(nums)
	k := 0
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] == nums[j-1] {
			k++
		}
		j--

	}

	if k == len(nums)/2 {
		answer = true
		fmt.Println(answer)
	} else {
		fmt.Println(answer)
	}
}
