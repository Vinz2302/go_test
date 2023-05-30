package handler

import (
	"errors"
	"fmt"
	"net/http"
	config "rest-api/app/config"
	model "rest-api/modules/v1/utilities/user/models"
	service "rest-api/modules/v1/utilities/user/service"
	res "rest-api/pkg/api-response"
	helper "rest-api/pkg/helpers"
	jwt "rest-api/pkg/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	conf, _ = config.Init()
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

	fmt.Println("email1 = ", LoginRequest.Email)
	refresh_token, errToken := jwt.RefreshToken(LoginRequest.Email)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("bad request"))
		return
	}

	fmt.Println("refresh token = ", refresh_token)

	LoginResult, err := h.userService.Login(LoginRequest, refresh_token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}

	errCompare := helper.CompareHash([]byte(LoginResult.Password), []byte(LoginRequest.Password))
	if errCompare != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("Invalid credentials"))
		return
	}

	token, err := jwt.GenerateToken(LoginResult.RoleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.ServerError(err.Error()))
		return
	}

	fmt.Println("token", token)
	c.SetCookie("Authorization", token, 3600*24*1, "", "", false, true)
	c.SetCookie("refresh_token", refresh_token, 3600*24*30, "", "", false, true)

	LoginResponse := responseUser(*LoginResult)
	c.JSON(http.StatusOK, res.Success(LoginResponse))

}

func (h *UserHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, res.StatusOK("success"))
}

func (h *UserHandler) RefreshAccessToken(c *gin.Context) {

	cookie, errCookie := c.Cookie("refresh_token")
	if errCookie != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("Token not found"))
		return
	}

	validate_token, err := jwt.ValidateToken(cookie, conf.App.Secret_key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, res.UnAuthorized(err.Error()))
		return
	}

	email, err := jwt.ExtractTokenEmail(validate_token)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("Invalid email"))
		return
	}

	RefreshResult, err, statusCode := h.userService.Refresh(email, cookie)
	if err != nil {
		if statusCode == 404 {
			c.JSON(http.StatusNotFound, res.NotFound(err.Error()))
			return
		}
		if statusCode == 400 {
			c.JSON(http.StatusBadRequest, res.BadRequest(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}

	token, err := jwt.GenerateToken(RefreshResult.RoleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.ServerError(err.Error()))
		return
	}

	//RefreshResponse := responseRefresh(token)

	c.JSON(http.StatusOK, res.Success(token))
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

/* func responseRefresh(b model.RefreshResponse) model.RefreshResponse {

	refreshResponse := model.RefreshResponse{
		Access_token: b.Access_token,
	}
	return refreshResponse
} */
