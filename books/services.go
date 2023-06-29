package books

type BookServiceInterface interface {
	CreateBook(book *Book) error
	UpdateBook(book *Book) error
	DeleteBook(bookId uint) error
	GetAllBooks() ([]Book, error)
	GetBookById(bookId uint) (*Book, error)
}

type BookServiceV1 struct {
	bookRepos BookReposInterface
}

func NewBookService() BookServiceInterface {
	return BookServiceV1{bookRepos: NewBookRepos()}
}

func (b BookServiceV1) CreateBook(book *Book) error {
	return b.bookRepos.CreateBook(book)
}

func (b BookServiceV1) UpdateBook(book *Book) error {
	return b.bookRepos.UpdateBook(book)
}

func (b BookServiceV1) DeleteBook(bookId uint) error {
	return b.bookRepos.DeleteBook(bookId)
}

func (b BookServiceV1) GetAllBooks() ([]Book, error) {
	return b.bookRepos.GetAllBooks()
}

func (b BookServiceV1) GetBookById(bookId uint) (*Book, error) {
	return b.bookRepos.GetBookById(bookId)
}
