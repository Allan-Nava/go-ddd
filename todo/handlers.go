package todo

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */
import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Allan-Nava/go-ddd/utils"
)

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

	todos, err := h.Service.GetAll()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(&utils.ApiError{Message: "not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(&utils.ApiError{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(todos)
}

///

// CreateTodo godoc
// @Summary      CreateTodo
// @Description  CreateTodo
// @Tags         todo
// @Accept json
// @Param customername body createTodoRequest true "customername"
// @Produce      json
// @Success      201
// @Router       /customers [post]
func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {

	requestBody := new(createTodoRequest)

	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}
	logrus.Debugf("into CreateCustomer with body:", requestBody)
	//
	errors := ValidateStruct(*requestBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}
	cust, err := h.Service.Create(c.Context(), requestBody.Name)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&utils.ApiError{Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(cust)
}
