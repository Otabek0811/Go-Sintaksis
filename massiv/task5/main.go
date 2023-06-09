package main

import "fmt"

func main() {
	var nums = []int{2, 3, 5, 4, 7}
	k := 3
	for _, val := range nums {
		for _, val2 := range nums {
			if val%val2 == k {
				fmt.Printf("(%d,%d) ", val, val2)
			}
		}
	}
	fmt.Println()
}
