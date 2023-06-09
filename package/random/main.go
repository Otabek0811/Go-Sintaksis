package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var nums = [500]int{}
	
	var  odd,even []int
	for index := range nums {
		a := rand.Intn(1000)-len(nums)/2
		nums[index] = a
	}
	for _,val := range nums {
		if val%2!=0{
			odd = append(odd, val)
		}else{
			even = append(even, val)
		}
	}
	fmt.Println(even)

	
}
