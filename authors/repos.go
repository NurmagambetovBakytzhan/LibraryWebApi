package authors

import (
	"Library_WebApi/books"
	"Library_WebApi/src"
	"github.com/jmoiron/sqlx"
)

type AuthorReposInterface interface {
	CreateAuthor(author *Author) error
	UpdateAuthor(author *Author) error
	DeleteAuthor(authorId uint) error
	GetAllAuthors() ([]Author, error)
	GetAuthor(authorId uint) (*Author, error)
	GetAvailableBooksByAuthorId(authorId uint) ([]books.Book, error)
}

type AuthorReposV1 struct {
	DB *sqlx.DB
}

//func NewAuthorRepos(db *sqlx.DB) AuthorReposInterface {
//	//db, _ := src.DbSetup()
//	return AuthorReposV1{DB: db}
//}

func NewAuthorRepos() AuthorReposInterface {
	db, _ := src.DbSetup()
	return &AuthorReposV1{DB: db}
}

func (a *AuthorReposV1) GetAvailableBooksByAuthorId(authorId uint) ([]books.Book, error) {
	var availableBooks []books.Book
	err := a.DB.Select(&availableBooks, `SELECT b.* FROM books b LEFT JOIN subscriptions s ON b.id = s.book_id WHERE b.author_id = $1 AND s.id IS NULL`, authorId)
	return availableBooks, err
}

func (a *AuthorReposV1) CreateAuthor(author *Author) error {
	_, err := a.DB.Exec("INSERT INTO authors (fio, pseudonym, specialization) VALUES ($1, $2, $3)", author.FIO, author.Pseudonym, author.Specialization)
	return err
}

func (a *AuthorReposV1) UpdateAuthor(author *Author) error {
	_, err := a.DB.Exec("UPDATE authors SET fio=$1, pseudonym=$2, specialization=$3 WHERE id=$4", author.FIO, author.Pseudonym, author.Specialization, author.ID)
	return err
}

func (a *AuthorReposV1) DeleteAuthor(authorId uint) error {
	_, err := a.DB.Exec("DELETE FROM authors WHERE id=$1", authorId)
	return err
}

func (a *AuthorReposV1) GetAllAuthors() ([]Author, error) {
	var authors []Author
	err := a.DB.Select(&authors, "SELECT * FROM authors")
	return authors, err
}

func (a *AuthorReposV1) GetAuthor(authorId uint) (*Author, error) {
	var author Author
	err := a.DB.Get(&author, "SELECT * FROM authors WHERE id=$1", authorId)
	if err != nil {
		return nil, err
	}
	return &author, nil
}
