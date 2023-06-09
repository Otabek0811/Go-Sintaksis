package main

import "fmt"

func main() {
	var nums = []int{2, 2, 2, 2, 1, 3, 1, 4}
	k := 2
	y := 0
	for _, val := range nums {
		x := 0
		for _, value := range nums {
			if val == value {
				x++
			}
		}
		if x == 1 {
			y++
		}
		if y == k {
			fmt.Println(val)
			break
		}

	}
	if y == 0 {
		fmt.Println(-1)
	}
}
