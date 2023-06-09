package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	y := 100
	var (
		random_number = rand.Intn(y)
		input         int
		chance        int = 0
	)
	k := 1
	for {
		k *= 2
		
		if k > y {
			break
		}
		chance++
	}

	//  fmt.Println("random_number:", random_number)

	for {
		fmt.Printf("Input number: ")
		fmt.Scan(&input)

		if input == random_number {
			fmt.Println("Find number 🎉🎉🎉")
			break
		} else {
			chance--
			if input < random_number {
				fmt.Println("Higher")
			} else {
				fmt.Println("Lower")
			}

		}

		fmt.Println("Chance: ", chance)

		if chance <= 0 {
			fmt.Println("You lose 😂")
			break
		}
	}
}

// 1. Top or Botttom
// 2. Limit => chance
