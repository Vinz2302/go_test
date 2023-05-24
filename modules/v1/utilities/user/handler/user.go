package handler

import (
	"errors"
	"fmt"
	"net/http"
	model "rest-api/modules/v1/utilities/user/models"
	service "rest-api/modules/v1/utilities/user/service"
	res "rest-api/pkg/api-response"
	helper "rest-api/pkg/helpers"
	jwt "rest-api/pkg/jwt"
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
	var LoginRequest model.UserLogin

	err := c.ShouldBindJSON(&LoginRequest)
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

	LoginResult, err := h.userService.Login(LoginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
	}

	errCompare := helper.CompareHash([]byte(LoginResult.Password), []byte(LoginRequest.Password))
	if errCompare != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("Invalid credentials"))
		return
	}

	token, err := jwt.GenerateToken(LoginResult.RoleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.ServerError(err.Error()))
	}

	fmt.Println("token", token)
	c.SetCookie("Authorization", token, 3600*24*1, "", "", false, true)

	LoginResponse := responseUser(*LoginResult)
	c.JSON(http.StatusOK, res.Success(LoginResponse))

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
