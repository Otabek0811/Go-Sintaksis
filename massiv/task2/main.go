package main

import "fmt"

func main() {
	var nums = []int{1, 1, 4, 5, 3, 3, 1, 3, 3, 1, 2}

	var nums2 = []int{}

	m := make(map[int]int)

	for _, val := range nums {
		k := 0
		for _, val2 := range nums2 {
			if val == val2 {
				k++
			}
		}
		if k == 0 {
			nums2 = append(nums2, val)
		}
	}

	for _, val2 := range nums2 {
		k := 0
		for _, val := range nums {
			if val == val2 {
				k++
			}
		}
		m[val2] = k
	}
	fmt.Println(m)
	max := -1

	f := 0
	for key, value := range m {
		if value > max {
			max = value
			f = key
		}
	}
	min := f
	for key, value := range m {
		if key < min {
			f = value
		}
	}

	fmt.Println(m[f])
}
