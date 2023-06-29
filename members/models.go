package members

import (
	"Library_WebApi/books"
)

type Member struct {
	ID    uint         `db:"id"`
	FIO   string       `db:"FIO"`
	Books []books.Book `db:"books"`
}
