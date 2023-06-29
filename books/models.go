package books

type Book struct {
	ID       uint   `db:"id"`
	Title    string `db:"title"`
	Genre    string `db:"genre"`
	ISBN     string `db:"ISBN"`
	AuthorId uint   `db:"author_id"`
}

type Subscription struct {
	ID       uint `db:"id"`
	BookId   uint `db:"bookId"`
	MemberId uint `db:"member_id"`
}
