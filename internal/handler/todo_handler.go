package handler

import (
	"strconv"

	"github.com/iskhakmuhamad/todo-api/internal/domain"
	"github.com/iskhakmuhamad/todo-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req domain.CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	todo, err := h.todoService.Create(userID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
		"data":    todo,
	})
}

func (h *TodoHandler) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	// Parse query parameters for filtering
	filter := domain.TodoFilter{
		Status:   domain.Status(c.Query("status")),
		Priority: domain.Priority(c.Query("priority")),
		Keyword:  c.Query("keyword"),
	}

	if categoryID := c.Query("category_id"); categoryID != "" {
		if id, err := strconv.ParseUint(categoryID, 10, 32); err == nil {
			filter.CategoryID = uint(id)
		}
	}

	if page := c.Query("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			filter.Page = p
		}
	}

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Limit = l
		}
	}

	todos, total, err := h.todoService.GetAll(userID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todos retrieved successfully",
		"data":    todos,
		"meta": fiber.Map{
			"total": total,
			"page":  filter.Page,
			"limit": filter.Limit,
		},
	})
}

func (h *TodoHandler) GetByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}

	todo, err := h.todoService.GetByID(uint(id), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo retrieved successfully",
		"data":    todo,
	})
}

func (h *TodoHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}

	var req domain.UpdateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	todo, err := h.todoService.Update(uint(id), userID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo updated successfully",
		"data":    todo,
	})
}

func (h *TodoHandler) Delete(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}

	if err := h.todoService.Delete(uint(id), userID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})
}

func (h *TodoHandler) ToggleStatus(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}

	todo, err := h.todoService.GetByID(uint(id), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	// Toggle status
	newStatus := domain.StatusTodo
	if todo.Status == domain.StatusTodo {
		newStatus = domain.StatusDone
	}

	updateReq := domain.UpdateTodoRequest{
		Status: newStatus,
	}

	updatedTodo, err := h.todoService.Update(uint(id), userID, updateReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo status toggled successfully",
		"data":    updatedTodo,
	})
}
