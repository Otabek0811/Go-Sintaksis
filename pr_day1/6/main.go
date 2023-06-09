package main

import "fmt"

func main(){
	var x int = 123

	d:=x%10
	x/=10
	c:=x%10
	x/=10

	fmt.Println(x*100+d*10+c)
}