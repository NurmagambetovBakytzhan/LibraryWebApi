package authors

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthorController struct {
	authorService AuthorServiceInterface
}

func NewAuthorController() AuthorController {
	return AuthorController{authorService: NewAuthorService()}
}

func (ac AuthorController) CreateAuthor(c *fiber.Ctx) error {
	var author Author
	if err := json.Unmarshal(c.Body(), &author); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	if err := ac.authorService.CreateAuthor(&author); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"Data": author})
}

func (ac *AuthorController) UpdateAuthor(c *fiber.Ctx) error {
	var author Author
	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	err := ac.authorService.UpdateAuthor(&author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(author)
}

func (ac *AuthorController) DeleteAuthor(c *fiber.Ctx) error {
	authorId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse id"})
	}

	err = ac.authorService.DeleteAuthor(uint(authorId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (ac *AuthorController) GetAllAuthors(c *fiber.Ctx) error {
	authors, err := ac.authorService.GetAllAuthors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(authors)
}

func (ac *AuthorController) GetAuthor(c *fiber.Ctx) error {
	authorId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse id"})
	}
	author, err := ac.authorService.GetAuthor(uint(authorId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(author)
}

func (ac *AuthorController) GetAvailableBooksByAuthorId(c *fiber.Ctx) error {
	authorId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse id"})
	}

	books, err := ac.authorService.GetAvailableBooksByAuthorId(uint(authorId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(books)
}
