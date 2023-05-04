package handler

import (
	"errors"
	"fmt"
	"net/http"
	model "rest-api/modules/v1/utilities/booking/models"
	service "rest-api/modules/v1/utilities/booking/services"
	res "rest-api/pkg/api-response"
	helper "rest-api/pkg/helpers"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookingHandler struct {
	bookingService service.IBookingService
}

func NewBookingHandler(bookingService service.IBookingService) *BookingHandler {
	return &BookingHandler{(bookingService)}
}

func (h *BookingHandler) Index(c *gin.Context) {

	limitString := c.Query("limit")
	pageString := c.Query("page")

	limit, errLimit := strconv.Atoi(limitString)
	if errLimit != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("limit required number"))
		return
	}

	page, errPage := strconv.Atoi(pageString)
	if errPage != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("page required number"))
		return
	}

	booking, count, err := h.bookingService.FindAll(page, limit)

	if len(booking) == 0 {
		c.JSON(http.StatusOK, res.Success(nil))
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	var bookingsResponse []model.BookingResponse
	for _, b := range booking {
		bookingResponse := responseBooking(b)
		bookingsResponse = append(bookingsResponse, bookingResponse)
	}

	endpoint := "v1/booking?"
	metadata := helper.PaginationMetadata(count, limit, &page, endpoint)
	result := res.Pagination{
		MetaData: &metadata,
		Records:  bookingsResponse,
	}
	c.JSON(http.StatusOK, res.Success(result))
}

func (h *BookingHandler) FindByID(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	BookingResult, err := h.bookingService.FindByID(id)

	//log.Print("time", BookingResult.StartTime)
	//log.Print("time", BookingResult.StartTime.Format("2006-01-02"))
	//log.Print("str", BookingResult.StartTimeStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	if BookingResult.ID == 0 {
		c.JSON(http.StatusNotFound, res.NotFound("ID"))
		return
	}

	BookingResponse := responseBooking(BookingResult)

	c.JSON(http.StatusOK, res.Success(BookingResponse))
}

func (h *BookingHandler) Create(c *gin.Context) {
	var BookingRequest model.BookingRequest
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	err := c.ShouldBindJSON(&BookingRequest)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorMessage := helper.ErrorMessage(err)
			c.JSON(http.StatusBadRequest, res.BadRequest(errorMessage))
			return
		}
		c.JSON(http.StatusBadRequest, res.BadRequest(err.Error()))
		return
	}

	/* if BookingRequest.BooktypeID > 2 {
		c.JSON(http.StatusNotFound, res.NotFound(err.Error()))
		return
	} */
	if today.After(BookingRequest.EndTime) {
		c.JSON(http.StatusBadRequest, res.BadRequest("End time must be further in the future"))
		return
	}

	if BookingRequest.BooktypeID == model.Driver {
		if BookingRequest.DriversID == nil {
			c.JSON(http.StatusBadRequest, res.BadRequest("Driver Required"))
			return
		}
	}

	if BookingRequest.BooktypeID != model.Driver && BookingRequest.BooktypeID != model.NonDriver {
		c.JSON(http.StatusBadRequest, res.BadRequest("Booktype id not exist"))
		return
	}
	fmt.Println("request = ", BookingRequest)

	BookingResult, err, statusCode := h.bookingService.Create(BookingRequest)
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

	fmt.Print("bookingresult", BookingResult)

	BookingResponse := responseBooking(*BookingResult)

	fmt.Print("BookingResponse = ", BookingResponse)

	c.JSON(http.StatusOK, res.Success(BookingResponse))
}

func (h *BookingHandler) Update(c *gin.Context) {
	var BookingRequest model.BookingRequest

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := c.ShouldBindJSON(&BookingRequest)
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

	//now := time.Now().UTC()
	//today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	/* if BookingRequest.StartTime.After(BookingRequest.EndTime) {
		c.JSON(http.StatusBadRequest, res.BadRequest("Start time must before end time"))
		return
	} */

	if BookingRequest.BooktypeID == model.Driver {
		if BookingRequest.DriversID == nil {
			c.JSON(http.StatusBadRequest, res.BadRequest("Driver ID required"))
			return
		}
	}

	if BookingRequest.BooktypeID == model.NonDriver {
		BookingRequest.DriversID = nil
	}

	if BookingRequest.BooktypeID != model.Driver && BookingRequest.BooktypeID != model.NonDriver {
		c.JSON(http.StatusBadRequest, res.BadRequest("Booktype id not exist"))
		return
	}
	fmt.Println("request = ", BookingRequest)

	BookingResult, err, statusCode := h.bookingService.Update(id, BookingRequest)
	if err != nil {
		if statusCode == 404 {
			c.JSON(http.StatusNotFound, res.NotFound("Customer ID"))
			return
		}
		if statusCode == 400 {
			c.JSON(http.StatusBadRequest, res.BadRequest(err.Error()))
			return
		}
		if statusCode == 403 {
			c.JSON(http.StatusForbidden, res.StatusForbidden(err.Error()))
		}
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
		return
	}
	BookingResponse := responseBooking(*BookingResult)
	//fmt.Print("bookingresult = ", *BookingResult)
	//fmt.Print("BookingResponse = ", BookingResponse)
	c.JSON(http.StatusOK, res.Success(BookingResponse))
}

func (h *BookingHandler) Delete(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err, statusCode := h.bookingService.Delete(id)
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

	c.JSON(http.StatusOK, res.StatusOK("Booking deleted Success"))
}

func (h *BookingHandler) Finish(c *gin.Context) {
	//var BookingRequest model.BookingRequest
	var FinishRequest model.FinishRequest

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := c.ShouldBindJSON(&FinishRequest)
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

	BookingResult, err, statusCode := h.bookingService.Finish(id, FinishRequest)
	if err != nil {
		if statusCode == 404 {
			c.JSON(http.StatusNotFound, res.NotFound("ID"))
			return
		}
		if statusCode == 400 {
			c.JSON(http.StatusBadRequest, res.BadRequest(err.Error()))
			return
		}
		if statusCode == 403 {
			c.JSON(http.StatusForbidden, res.StatusForbidden(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, res.ServerError(err.Error()))
	}

	BookingResponse := responseBooking(*BookingResult)

	c.JSON(http.StatusOK, res.Success(BookingResponse))
	/* err, statusCode := h.bookingService.Finish(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}
	if statusCode == 404 {
		c.JSON(http.StatusNotFound, res.NotFound("ID"))
		return
	}

	c.JSON(http.StatusOK, res.StatusOK("Booking finished")) */
}

func responseBooking(b model.Booking) model.BookingResponse {

	bookingResponse := model.BookingResponse{
		ID:              b.ID,
		CustomersID:     b.CustomersID,
		CarsID:          b.CarsID,
		StartTime:       b.StartTime,
		EndTime:         b.EndTime,
		DriversID:       b.DriversID,
		TotalCost:       b.TotalCost,
		Finished:        b.Finished,
		Discount:        b.Discount,
		TotalDriverCost: b.TotalDriverCost,

		Booktype: b.BooktypeID,
	}
	return bookingResponse
}
