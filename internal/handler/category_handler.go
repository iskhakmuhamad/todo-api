package handler

import (
	"strconv"

	"github.com/iskhakmuhamad/todo-api/internal/domain"
	"github.com/iskhakmuhamad/todo-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req domain.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	category, err := h.categoryService.Create(userID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category created successfully",
		"data":    category,
	})
}

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	categories, err := h.categoryService.GetAll(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Categories retrieved successfully",
		"data":    categories,
	})
}

func (h *CategoryHandler) GetByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	category, err := h.categoryService.GetByID(uint(id), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Category not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category retrieved successfully",
		"data":    category,
	})
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	var req domain.UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	category, err := h.categoryService.Update(uint(id), userID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category updated successfully",
		"data":    category,
	})
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	if err := h.categoryService.Delete(uint(id), userID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}
