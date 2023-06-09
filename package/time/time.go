package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	name      string
	price     int
	createdAt string
}

func generateOrder(count int) []Order {

	var orders []Order
	for i := 1; i <= count; i++ {
		orders = append(orders, Order{
			id:        i,
			name:      "apple_" + fmt.Sprintf("%d", i),
			price:     i * 100,
			createdAt: time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return orders
}

func updateTime(n string) string {
	date, err := time.Parse("2006-01-02 15:04:05", n)
	if err != nil {
		fmt.Println(err)
	}

	return date.Add(time.Hour * 2).Format("2006-01-02 15:04:05")
}

func main() {

	var orders = generateOrder(10)

	for _, order := range orders {

		order.createdAt = updateTime(order.createdAt)
		// fmt.Printf("%+v\n", order)
	}

	currentTime := time.Now()
	startOfYear := time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC)
	result := currentTime.Sub(startOfYear)
	fmt.Println(result) // 3488h22m10.4962529s
	// fmt.Println(time.Now())
	// fmt.Println(time.Now().Add(time.Hour * 10))

	// time.Sleep(4 * time.Second)

	// fmt.Println("Finished...")

	// nowTime := time.Now()
	// created := time.Now().Add(time.Hour * -2)

	// fmt.Println(nowTime.After(created))
	// fmt.Println(created.After(nowTime))

	// fmt.Println(nowTime.Before(created))
	// fmt.Println(created.Before(nowTime))
}
