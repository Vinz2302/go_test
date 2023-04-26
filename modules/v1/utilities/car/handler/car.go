package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	model "rest-api/modules/v1/utilities/car/models"
	service "rest-api/modules/v1/utilities/car/services"
	res "rest-api/pkg/api-response"
	helper "rest-api/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CarHandler struct {
	carService service.ICarService
}

func NewCarHandler(carService service.ICarService) *CarHandler {
	return &CarHandler{carService}
}

func (h *CarHandler) Index(c *gin.Context) {
	limitString := c.Query("limit")
	pageString := c.Query("page")

	limit, errLimit := strconv.Atoi(limitString)
	if errLimit != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("Limit required number"))
		return
	}

	page, errPage := strconv.Atoi(pageString)
	if errPage != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("Page required number"))
		return
	}

	car, count, err := h.carService.FindAll(page, limit)

	if len(car) == 0 {
		c.JSON(http.StatusOK, res.Success(nil))
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	var carsResponse []model.CarResponse
	for _, b := range car {
		carResponse := responseCar(b)
		carsResponse = append(carsResponse, carResponse)
	}

	endpoint := "v1/car?"
	metadata := helper.PaginationMetadata(count, limit, &page, endpoint)
	result := res.Pagination{
		MetaData: &metadata,
		Records:  carsResponse,
	}
	c.JSON(http.StatusOK, res.Success(result))
}

func (h *CarHandler) FindByID(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	CarResult, err := h.carService.FindByID(id)

	log.Print(CarResult)

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	if CarResult.ID == 0 {
		c.JSON(http.StatusNotFound, res.NotFound("ID"))
		return
	}

	CarResponse := responseCar(CarResult)

	c.JSON(http.StatusOK, res.Success(CarResponse))
}

func (h *CarHandler) Create(c *gin.Context) {
	var CarRequest model.CarRequest

	err := c.ShouldBindJSON(&CarRequest)
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

	CarResult, err := h.carService.Create(CarRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}
	CarResponse := responseCar(*CarResult)
	c.JSON(http.StatusOK, res.Success(CarResponse))
}

func (h *CarHandler) Update(c *gin.Context) {
	var CarRequest model.CarRequest

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := c.ShouldBindJSON(&CarRequest)
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

	CarResult, err, statusCode := h.carService.Update(id, CarRequest)
	if err != nil {
		if statusCode == 404 {
			c.JSON(http.StatusNotFound, res.NotFound("ID"))
			return
		}
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}
	CarResponse := responseCar(*CarResult)
	c.JSON(http.StatusOK, res.Success(CarResponse))
}

func (h *CarHandler) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	fmt.Println(id)

	err, statusCode := h.carService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}

	if statusCode == 404 {
		c.JSON(http.StatusNotFound, res.NotFound("ID"))
		return
	}
	c.JSON(http.StatusOK, res.StatusOK("Car deleted Success"))
}

func responseCar(b model.Car) model.CarResponse {

	carResponse := model.CarResponse{
		ID:             b.ID,
		CarName:        b.CarName,
		RentDailyPrice: b.RentDailyPrice,
		Stock:          b.Stock,
	}
	return carResponse
}
