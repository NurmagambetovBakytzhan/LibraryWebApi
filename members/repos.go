package members

import (
	"Library_WebApi/books"
	"Library_WebApi/src"
	"github.com/jmoiron/sqlx"
)

type MembersReposInterface interface {
	CreateMember(member *Member) error
	UpdateMember(member *Member) error
	DeleteMember(memberId uint) error
	GetAllMembers() ([]Member, error)
	GetMemberById(memberId uint) (*Member, error)
	GetBooksByMemberId(memberId uint) ([]books.Book, error)
}

type MemberReposV1 struct {
	DB *sqlx.DB
}

func NewMemberRepos() MembersReposInterface {
	db, _ := src.DbSetup()
	return &MemberReposV1{DB: db}
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
	err := m.DB.Select(&books, `SELECT b.* FROM books b LEFT JOIN subscriptions s ON b.id = s.bookId WHERE s.member_id = $1`, memberId)
	return books, err
}
