package main

import "fmt"
func main(){
	var x, y, z int = 12,22,3

	min:=x
	if min>y{
		min=y
	}
	if min>z{
		min=z
	}

	fmt.Println(min)

}