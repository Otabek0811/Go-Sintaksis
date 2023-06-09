package main

type bookCollection struct{
	books []*Book
}

func (b *bookCollection) createIterator()Iterator_Book{
	return &BookIterator{
		books: b.books,
	}
}