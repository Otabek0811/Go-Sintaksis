package main

import "fmt"

func main() {
	var nums = []int{10, 12, 11, 15}
	min := nums[0]
	max := nums[0]
	for _, val := range nums {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Println("Low:", min, "High:", max)

	for {
		k := 0
		for _, val := range nums {
			if min == val {
				k++
			}

		}
		if k == 0 {
			fmt.Println(min)
		}
		min++
		if min == max {
			break
		}
	}

}
