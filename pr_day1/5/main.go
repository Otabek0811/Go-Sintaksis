package main

import "fmt"

func main() {
	var x int = 656
	var y int
	d := x
	summ := 0
	for{
		y = x%10
		summ = summ*10 + y
		x /= 10

		if(x==0){
			break 
		}
	}

	if summ == d {
		fmt.Println("Palindrome")
	} else {
		fmt.Println(summ, "Not Palindrome")
	}

}
