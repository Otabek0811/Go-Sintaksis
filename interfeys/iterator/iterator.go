package main

type Iterator interface {
	hasNext() bool
	getNext() *User
}

type Iterator_Book interface{
	hasNext() bool
	getNext() *Book
}
