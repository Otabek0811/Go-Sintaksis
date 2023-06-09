package main

import "fmt"

func main() {
	user1 := &User{
		name: "Asliddin",
		age:  21,
	}

	user2 := &User{
		name: "Tursinbek",
		age:  19,
	}

	userCollection := &userCollection{
		users: []*User{user1, user2},
	}
	Iterator := userCollection.createIterator()

	for Iterator.hasNext() {
		user := Iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}

	book1:=&Book{
		name: "Otkan Kunlar",
		price: 5,
	}

	book2:=&Book{
		name: "Qaytganimda Uyda Bol",
		price: 6,
	}
	bookCollection:=&bookCollection{
		books: []*Book{book1,book2},
	}

	BookIterator:=bookCollection.createIterator()

	if BookIterator.hasNext(){
		book:=BookIterator.getNext()
		fmt.Printf("Book is %+v\n", book)
	}
}
