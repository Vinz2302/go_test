package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	model "rest-api/modules/v1/utilities/drivers/models"
	service "rest-api/modules/v1/utilities/drivers/services"
	"strconv"

	res "rest-api/pkg/api-response"
	helper "rest-api/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DriverHandler struct {
	driverService service.IDriverService
}

func NewDriverHandler(driverService service.IDriverService) *DriverHandler {
	return &DriverHandler{driverService}
}

func (h *DriverHandler) Index(c *gin.Context) {
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

	driver, count, err := h.driverService.FindAll(page, limit)

	if len(driver) == 0 {
		c.JSON(http.StatusOK, res.Success(nil))
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	var driversResponse []model.DriverResponse
	for _, b := range driver {
		driverResponse := responseDriver(b)
		driversResponse = append(driversResponse, driverResponse)
	}
	endpoint := "v1/driver?"
	metadata := helper.PaginationMetadata(count, limit, &page, endpoint)
	result := res.Pagination{
		MetaData: &metadata,
		Records:  driversResponse,
	}
	c.JSON(http.StatusOK, res.Success(result))
}

func (h *DriverHandler) FindByID(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	DriverResult, err := h.driverService.FindByID(id)

	log.Print(DriverResult)

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	if DriverResult.ID == 0 {
		c.JSON(http.StatusNotFound, res.NotFound("ID"))
		return
	}

	DriverResponse := responseDriver(DriverResult)

	c.JSON(http.StatusOK, res.Success(DriverResponse))
}

func (h *DriverHandler) Create(c *gin.Context) {
	var DriverRequest model.DriverRequest

	err := c.ShouldBindJSON(&DriverRequest)
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

	DriverResult, err := h.driverService.Create(DriverRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}
	DriverResponse := responseDriver(*DriverResult)
	c.JSON(http.StatusOK, res.Success(DriverResponse))
}

func (h *DriverHandler) Update(c *gin.Context) {
	var DriverRequest model.DriverRequest

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := c.ShouldBindJSON(&DriverRequest)
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

	DriverResult, err, statusCode := h.driverService.Update(id, DriverRequest)
	if err != nil {
		if statusCode == 404 {
			c.JSON(http.StatusNotFound, res.NotFound("ID"))
			return
		}
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}
	DriverResponse := responseDriver(*DriverResult)
	c.JSON(http.StatusOK, res.Success(DriverResponse))
}

func (h *DriverHandler) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	fmt.Println(id)

	err, statusCode := h.driverService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}

	if statusCode == 404 {
		c.JSON(http.StatusNotFound, res.NotFound("ID"))
		return
	}
	c.JSON(http.StatusOK, res.StatusOK("Driver deleted Success"))
}

func responseDriver(b model.Driver) model.DriverResponse {

	driverResponse := model.DriverResponse{
		ID:          b.ID,
		DriverName:  b.DriverName,
		Nik:         b.Nik,
		PhoneNumber: b.PhoneNumber,
		DailyCost:   b.DailyCost,
	}
	return driverResponse
}
