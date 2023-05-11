package handler

import (
	"errors"
	"fmt"
	"net/http"
	model "rest-api/modules/v1/utilities/user/models"
	service "rest-api/modules/v1/utilities/user/service"
	res "rest-api/pkg/api-response"
	helper "rest-api/pkg/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	fmt.Println("id = ", id)

	UserResult, err := h.userService.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	if UserResult.ID == 0 {
		c.JSON(http.StatusNotFound, res.NotFound("ID"))
		return
	}

	UserResponse := responseUser(UserResult)

	c.JSON(http.StatusOK, res.Success(UserResponse))
}

func (h *UserHandler) Create(c *gin.Context) {
	var UserRequest model.UserRequest

	err := c.ShouldBindJSON(&UserRequest)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorMessages := helper.ErrorMessage(err)
			c.JSON(http.StatusBadRequest, res.BadRequest(errorMessages))
			return
		}
		c.JSON(http.StatusBadRequest, res.BadRequest(err.Error()))
		return
	}

	fmt.Println("user request = ", UserRequest)

	UserResult, err := h.userService.Create(UserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}
	UserResponse := responseUser(*UserResult)
	c.JSON(http.StatusOK, res.Success(UserResponse))

}

func (h *UserHandler) Login(c *gin.Context) {
	var UserLogin model.UserLogin

	err := c.ShouldBindJSON(&UserLogin)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorMessages := helper.ErrorMessage(err)
			c.JSON(http.StatusBadRequest, res.BadRequest(errorMessages))
			return
		}
		c.JSON(http.StatusBadRequest, res.BadRequest(err.Error()))
		return
	}

	//var user model.Users

}

func responseUser(b model.Users) model.UserResponse {

	userResponse := model.UserResponse{
		ID:     b.ID,
		Name:   b.Name,
		Email:  b.Email,
		RoleId: b.RoleId,
	}
	return userResponse
}
