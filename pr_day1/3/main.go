package main

import "fmt"

func main(){
	var a,b,c uint8=1,2,2

	if a==b && c==b {
		fmt.Println("Teng tomonli")
	}else if a==b || c==b || a==c{
		fmt.Println("Teng Yonli")
	}else{
		fmt.Println("Turli tomonli")
	}
}