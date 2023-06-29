package members

import (
	"Library_WebApi/books"
)

type MemberServiceInterface interface {
	CreateMember(member *Member) error
	UpdateMember(member *Member) error
	DeleteMember(memberId uint) error
	GetAllMembers() ([]Member, error)
	GetMemberById(memberId uint) (*Member, error)
	GetBooksByMemberId(memberId uint) ([]books.Book, error)
	SubscribeBookById(memberId uint, bookId uint) error
	UnsubscribeBookById(memberId uint, bookId uint) error
}

type MemberServiceV1 struct {
	memberRepos MembersReposInterface
}

func NewMemberService() MemberServiceInterface {
	return MemberServiceV1{memberRepos: NewMemberRepos()}
}

func (m MemberServiceV1) SubscribeBookById(memberId uint, bookId uint) error {
	return m.memberRepos.SubscribeBookById(memberId, bookId)
}
func (m MemberServiceV1) UnsubscribeBookById(memberId uint, bookId uint) error {
	return m.memberRepos.UnsubscribeBookById(memberId, bookId)
}

func (m MemberServiceV1) CreateMember(member *Member) error {
	return m.memberRepos.CreateMember(member)
}

func (m MemberServiceV1) UpdateMember(member *Member) error {
	return m.memberRepos.UpdateMember(member)
}

func (m MemberServiceV1) DeleteMember(memberId uint) error {
	return m.memberRepos.DeleteMember(memberId)
}

func (m MemberServiceV1) GetAllMembers() ([]Member, error) {
	return m.memberRepos.GetAllMembers()
}

func (m MemberServiceV1) GetMemberById(memberId uint) (*Member, error) {
	return m.memberRepos.GetMemberById(memberId)
}

func (m MemberServiceV1) GetBooksByMemberId(memberId uint) ([]books.Book, error) {
	return m.memberRepos.GetBooksByMemberId(memberId)
}
