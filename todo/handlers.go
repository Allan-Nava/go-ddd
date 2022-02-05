package todo

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Allan-Nava/go-ddd/utils"
)

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */

type TodoHandler struct {
	Service *TodoService
}

// GetAll godoc
// @Summary      Get all todo
// @Description  get all todo
// @Tags         todo
// @Produce      json
// @Success      200  {array}  CustomerCompleteInfo
// @Router       /todos [get]
func (h *TodoHandler) GetAll(c *fiber.Ctx) error {
	logrus.Debugf("into GetAll")

	todos, err := h.Service.All()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(&utils.ApiError{Message: "not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(&utils.ApiError{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(todos)
}

///
