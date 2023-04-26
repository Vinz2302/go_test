package handler

import (
	"fmt"
	"net/http"
	model "rest-api/modules/v1/utilities/report/models"
	service "rest-api/modules/v1/utilities/report/services"
	res "rest-api/pkg/api-response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportService service.IReportService
}

func NewReportHandler(reportService service.IReportService) *ReportHandler {
	return &ReportHandler{(reportService)}
}

func (h *ReportHandler) FindMonthlyCompanyIncome(c *gin.Context) {
	yearString := c.Query("year")
	monthString := c.Query("month")

	year, errYear := strconv.Atoi(yearString)
	if errYear != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("year required number"))
		return
	}

	month, errMonth := strconv.Atoi(monthString)
	if errMonth != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("month required number"))
	}

	reportActivity, err := h.reportService.FindMonthlyCompanyIncome(year, month)

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	var reportsResponse []model.ReportResponse
	for _, b := range reportActivity {
		reportBooking := responseReport(b)
		reportsResponse = append(reportsResponse, reportBooking)
	}

	fmt.Println("reportresponse, ", reportsResponse)

	c.JSON(http.StatusOK, res.Success(reportsResponse))

}

func (h *ReportHandler) FindBookingActivity(c *gin.Context) {
	yearString := c.Query("year")
	monthString := c.Query("month")

	year, errYear := strconv.Atoi(yearString)
	if errYear != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("year required number"))
		return
	}

	month, errMonth := strconv.Atoi(monthString)
	if errMonth != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("month required number"))
		return
	}

	reportCar, err := h.reportService.FindBookingActivity(year, month)

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	//reportCarResponse := responseReportCar(reportCar)
	var carsResponse []model.ReportCar
	for _, b := range reportCar {
		reportCarResponse := responseReportCar(b)
		carsResponse = append(carsResponse, reportCarResponse)
	}

	fmt.Println("reportCar = ", carsResponse)

	c.JSON(http.StatusOK, res.Success(carsResponse))
}

func (h *ReportHandler) FindDriverActivity(c *gin.Context) {
	yearString := c.Query("year")
	monthstring := c.Query("month")

	year, errYear := strconv.Atoi(yearString)
	if errYear != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("year required number"))
		return
	}

	month, errMonth := strconv.Atoi(monthstring)
	if errMonth != nil {
		c.JSON(http.StatusBadRequest, res.BadRequest("month required number"))
		return
	}

	reportDriver, err := h.reportService.FindDriverActivity(year, month)

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ServerError(err))
		return
	}

	var driversResponse []model.ReportDriver
	for _, b := range reportDriver {
		reportDriverResponse := responseReportDriver(b)
		driversResponse = append(driversResponse, reportDriverResponse)
	}
	fmt.Println("rerportCar = ", driversResponse)

	c.JSON(http.StatusOK, res.Success(driversResponse))
}

func responseReport(b model.Report) model.ReportResponse {
	reportResponse := model.ReportResponse{
		BookingID:            b.BookingID,
		TotalDriverCost:      b.TotalDriverCost,
		TotalDriverIncentive: b.TotalDriverIncentive,
		TotalDriverExpense:   b.TotalDriverExpense,
		TotalGrossIncome:     b.TotalGrossIncome,
		TotalNettIncome:      b.TotalNettIncome,
	}
	return reportResponse
}

func responseReportCar(b model.ReportCar) model.ReportCar {
	reportCarResponse := model.ReportCar{
		CarID:             b.CarID,
		CarName:           b.CarName,
		TotalDaysBooking:  b.TotalDaysBooking,
		TotalBookingCount: b.TotalBookingCount,
	}
	return reportCarResponse
}

func responseReportDriver(b model.ReportDriver) model.ReportDriver {
	reportDriverResponse := model.ReportDriver{
		DriverID:         b.DriverID,
		DriverName:       b.DriverName,
		TotalDriverDay:   b.TotalDriverDay,
		TotalDriverCount: b.TotalDriverCount,
		TotalIncentive:   b.TotalIncentive,
	}
	return reportDriverResponse
}
