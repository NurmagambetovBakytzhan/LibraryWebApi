package members

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type MemberController struct {
	memberService MemberServiceInterface
}

func NewMemberController() MemberController {
	return MemberController{memberService: NewMemberService()}
}

func (mc MemberController) CreateMember(c *fiber.Ctx) error {
	var member Member
	if err := json.Unmarshal(c.Body(), &member); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := mc.memberService.CreateMember(&member); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"Data": member})
}

func (mc MemberController) UpdateMember(c *fiber.Ctx) error {
	var member Member
	if err := json.Unmarshal(c.Body(), &member); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	err := mc.memberService.UpdateMember(&member)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}

func (mc MemberController) DeleteMember(c *fiber.Ctx) error {
	var memberId uint
	if err := json.Unmarshal(c.Body(), &memberId); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	err := mc.memberService.DeleteMember(memberId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (mc MemberController) GetAllMembers(c *fiber.Ctx) error {
	members, err := mc.memberService.GetAllMembers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(members)
}

func (mc MemberController) GetMemberById(c *fiber.Ctx) error {
	var memberId uint
	if err := c.BodyParser(memberId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	member, err := mc.memberService.GetMemberById(memberId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}

func (mc MemberController) GetBooksByMemberId(c *fiber.Ctx) error {
	var memberId uint
	if err := c.BodyParser(memberId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	books, err := mc.memberService.GetBooksByMemberId(memberId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)

}