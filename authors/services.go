package authors

import "Library_WebApi/books"

type AuthorServiceInterface interface {
	CreateAuthor(author *Author) error
	UpdateAuthor(author *Author) error
	DeleteAuthor(authorId uint) error
	GetAllAuthors() ([]Author, error)
	GetAuthor(authorId uint) (*Author, error)
	GetAvailableBooksByAuthorId(authorId uint) ([]books.Book, error)
}

type AuthorServiceV1 struct {
	authorRepos AuthorReposInterface
}

func NewAuthorService() AuthorServiceInterface {
	return AuthorServiceV1{authorRepos: NewAuthorRepos()}
}

func (a AuthorServiceV1) CreateAuthor(author *Author) error {
	return a.authorRepos.CreateAuthor(author)
}
func (a AuthorServiceV1) UpdateAuthor(author *Author) error {
	return a.authorRepos.UpdateAuthor(author)
}
func (a AuthorServiceV1) DeleteAuthor(authorId uint) error {
	return a.authorRepos.DeleteAuthor(authorId)
}
func (a AuthorServiceV1) GetAllAuthors() ([]Author, error) {
	return a.authorRepos.GetAllAuthors()
}
func (a AuthorServiceV1) GetAuthor(authorId uint) (*Author, error) {
	return a.authorRepos.GetAuthor(authorId)
}
func (a AuthorServiceV1) GetAvailableBooksByAuthorId(authorId uint) ([]books.Book, error) {
	return a.authorRepos.GetAvailableBooksByAuthorId(authorId)
}
