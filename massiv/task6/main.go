package main

import "fmt"

func main() {
	var nums = []int{1,5,7,-1}
	sum:=6
	k := 0
	for idx, val := range nums {
		for i:=idx;i<len(nums);i++ {
			if val+nums[i] == sum {
				k++
			}
		}
	}
	fmt.Println(k)
}