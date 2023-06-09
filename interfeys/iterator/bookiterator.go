package main

type BookIterator struct {
	index int
	books []*Book
}

func (b *BookIterator) hasNext() bool {
	if b.index < len(b.books) {
		return true
	}
	return false
}

func (b *BookIterator) getNext() *Book {
	if b.hasNext() {
		book := b.books[b.index]
		b.index++
		return book
	}
	return nil
}
