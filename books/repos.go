package books

import (
	"Library_WebApi/src"
	"github.com/jmoiron/sqlx"
)

type BookReposInterface interface {
	CreateBook(book *Book) error
	UpdateBook(book *Book) error
	DeleteBook(bookId uint) error
	GetAllBooks() ([]Book, error)
	GetBookById(bookId uint) (*Book, error)
}

type BookReposV1 struct {
	DB *sqlx.DB
}

func NewBookRepos() (BookReposInterface) {
	db, _ := src.DbSetup()

	return &BookReposV1{DB: db}
}

func (b *BookReposV1) CreateBook(book *Book) error {
	_, err := b.DB.Exec("INSERT INTO books (title, genre, isbn, author_id) VALUES ($1, $2, $3, $4)", book.Title, book.Genre, book.ISBN, book.AuthorId)
	return err
}

func (b *BookReposV1) UpdateBook(book *Book) error {
	_, err := b.DB.Exec("UPDATE books SET title=$1, genre=$2, isbn=$3, author_id=$4 WHERE id=$5", book.Title, book.Genre, book.ISBN, book.AuthorId)
	return err
}

func (b *BookReposV1) DeleteBook(bookId uint) error {
	_, err := b.DB.Exec("DELETE FROM books WHERE id=$1", bookId)
	return err
}

func (b *BookReposV1) GetAllBooks() ([]Book, error) {
	var books []Book
	err := b.DB.Select(&books, "SELECT * FROM books")
	return books, err
}

func (b *BookReposV1) GetBookById(bookId uint) (*Book, error) {
	var book Book
	err := b.DB.Get(&book, "SELECT * FROM books WHERE id=$1", bookId)
	return &book, err
}
