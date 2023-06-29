package members

import (
	"Library_WebApi/books"
	"Library_WebApi/src"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type MembersReposInterface interface {
	CreateMember(member *Member) error
	UpdateMember(member *Member) error
	DeleteMember(memberId uint) error
	GetAllMembers() ([]Member, error)
	GetMemberById(memberId uint) (*Member, error)
	GetBooksByMemberId(memberId uint) ([]books.Book, error)
	SubscribeBookById(memberId uint, bookId uint) error
	UnsubscribeBookById(memberId uint, bookId uint) error
}

type MemberReposV1 struct {
	DB *sqlx.DB
}

func NewMemberRepos() MembersReposInterface {
	db, _ := src.DbSetup()
	return &MemberReposV1{DB: db}
}

func (m *MemberReposV1) UnsubscribeBookById(memberId uint, bookId uint) error {
	var count int
	err := m.DB.Get(&count, "SELECT COUNT(*) FROM books WHERE id = $1", bookId)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("The book does not exist")
	}

	err = m.DB.Get(&count, "SELECT count(*) from subscriptions WHERE book_id=$1 AND member_id=$2 AND deleted_at is null", bookId, memberId)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("The member is not subscribed to the book")
	}
	deleted_at := time.Now()
	_, err = m.DB.Exec("UPDATE subscriptions SET deleted_at=$1 WHERE book_id=$2 AND member_id=$3", deleted_at, bookId, memberId)
	if err != nil {
		return err
	}
	return nil
}

func (m *MemberReposV1) SubscribeBookById(memberId uint, bookId uint) error {
	var count int
	err := m.DB.Get(&count, "SELECT COUNT(*) FROM books WHERE id = $1", bookId)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("The book does not exist")
	}

	err = m.DB.Get(&count, "SELECT count(*) from subscriptions WHERE book_id=$1 AND deleted_at is null", bookId)
	if err != nil {
		return err
	}

	if count == 0 {
		fmt.Println("The book is already subscribed by someone")
		return err
	}

	createdAt := time.Now()
	_, err = m.DB.Exec("INSERT INTO subscriptions (book_id, member_id, created_at) VALUES ($1,$2,$3)", bookId, memberId, createdAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *MemberReposV1) CreateMember(member *Member) error {
	_, err := m.DB.Exec("INSERT INTO members (FIO) VALUES ($1)", member.FIO)
	return err
}

func (m *MemberReposV1) UpdateMember(member *Member) error {
	_, err := m.DB.Exec("UPDATE members SET FIO=$1 WHERE id=$2", member.FIO, member.ID)
	return err
}

func (m *MemberReposV1) DeleteMember(memberId uint) error {
	_, err := m.DB.Exec("DELETE FROM members WHERE id=$1", memberId)
	return err
}

func (m *MemberReposV1) GetAllMembers() ([]Member, error) {
	var members []Member
	err := m.DB.Select(&members, "SELECT * FROM members")
	return members, err
}

func (m *MemberReposV1) GetMemberById(memberId uint) (*Member, error) {
	var member Member
	err := m.DB.Get(&member, "SELECT * FROM members WHERE id=$1", memberId)
	return &member, err
}

func (m *MemberReposV1) GetBooksByMemberId(memberId uint) ([]books.Book, error) {
	var books []books.Book
	err := m.DB.Select(&books, `SELECT b.* FROM books b LEFT JOIN subscriptions s ON b.id = s.book_id WHERE s.member_id = $1`, memberId)
	return books, err
}
