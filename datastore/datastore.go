package datastore

import "github.com/isadoraora/API-GO/loader"

type BookStore interface {
	Initialize()
	SearchAuthor(author string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.BookData
	SearchBook(bookName string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.BookData
	SearchISBN(isbn string) *loader.BookData
	CreateBook(book *loader.BookData) bool
	DeleteBook(isbn string) bool
	UpdateBook(isbn string, book *loader.BookData) bool
}