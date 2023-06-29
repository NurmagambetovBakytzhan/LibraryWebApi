package books

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type BookController struct {
	bookService BookServiceInterface
}

func NewBookController() BookController {
	return BookController{bookService: NewBookService()}
}

func (bc BookController) CreateBook(c *fiber.Ctx) error {
	var book Book
	if err := json.Unmarshal(c.Body(), &book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := bc.bookService.CreateBook(&book); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"Data": book})
}

func (bc BookController) UpdateBook(c *fiber.Ctx) error {
	var book Book
	if err := json.Unmarshal(c.Body(), &book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	err := bc.bookService.UpdateBook(&book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(book)
}

func (bc BookController) DeleteBook(c *fiber.Ctx) error {
	var bookId uint
	if err := json.Unmarshal(c.Body(), &bookId); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	err := bc.bookService.DeleteBook(bookId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (bc BookController) GetAllBooks(c *fiber.Ctx) error {
	books, err := bc.bookService.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)
}

func (bc BookController) GetBookById(c *fiber.Ctx) error {
	var bookId uint
	if err := c.BodyParser(bookId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	book, err := bc.bookService.GetBookById(bookId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(book)
}
