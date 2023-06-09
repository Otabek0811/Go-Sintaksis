package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
}

// func BubbleSort(nums []int) []int {
// 	count := 0
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums)-i-1; j++ {
// 			if nums[j] > nums[j+1] {
// 				nums[j], nums[j+1] = nums[j+1], nums[j]
// 			}
// 			count++
// 		}
// 	}

// 	fmt.Println(count)

// 	return nums
// }

func main() {

	// var nums = []int{4, 2, 60, 23, 56, 1, 9, 6, 7}
	// var names = []string{"Ali", "Vali", "Asadbek", "Rahm"}
	// // nums = BubbleSort(nums)

	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })

	// sort.Slice(names, func(i, j int) bool {
	// 	return len(names[i]) < len(names[j])
	// })

	// fmt.Println(nums)
	// fmt.Println(names)

	var users = []User{
		{name: "Asadbek", age: 23},
		{name: "Mirjalol", age: 20},
		{name: "Asliddin", age: 21},
		{name: "Tursinbek", age: 18},
		{name: "Shohruh", age: 19},
		{name: "Otabek", age: 22},
		{name: "Doniyor", age: 21},
		{name: "Aziz", age: 25},
		{name: "Abdulbosit", age: 19},
	}

	// users = sorting(users, "AGE")
	users = sorting(users, "NAME")

	for _, user := range users {
		fmt.Println(user)
	}
}

func sorting(s []User, x string) []User {
	if x == "NAME" {
		sort.Slice(s, func(i, j int) bool {
			return s[i].name <s[j].name
		})
	} else if x == "AGE" {
		sort.Slice(s, func(i, j int) bool {
			return s[i].age<s[j].age
		})
	}
	return s
}
