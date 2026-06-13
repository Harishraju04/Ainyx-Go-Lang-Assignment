package handler

import (
	"net/http"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) ListAllUsers(c *fiber.Ctx) error {
	var paginationQuery PaginationQuery
	if err := c.QueryParser(&paginationQuery); err != nil {
		return handleError(c, err)
	}

	// Set defaults if not provided
	if paginationQuery.Page == 0 {
		paginationQuery.Page = 1
	}
	if paginationQuery.PageSize == 0 {
		paginationQuery.PageSize = 10
	}

	if err := validator.Validate.Struct(paginationQuery); err != nil {
		return handleError(c, err)
	}

	res, err := h.svc.ListAllUsersPaginated(c.Context(), paginationQuery.Page, paginationQuery.PageSize)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(res)
}
