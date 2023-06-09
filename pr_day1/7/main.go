package main

import "fmt"

func main(){
	var A,B,C int = 1,2,3

	A,B,C = B,C,A

	fmt.Println(A,B,C)
}