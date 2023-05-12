package helpers

import (
	"fmt"
	"math"
	"rest-api/app/config"
	helperDatabases "rest-api/pkg/helpers/databases"
	"time"
)

var (
	conf, err = config.Init()
)

func PaginationMetadata(count *int64, limit int, page *int, endpoint string) helperDatabases.ResponseBackPaginationDTO {

	var totalPage float64
	previousPage := *page - 1
	nextPage := *page + 1
	totalData := int(*count)
	totalPage1 := float64(totalData) / float64(limit)
	totalPage1 = math.Ceil(totalPage1)

	if totalPage1 == 0 {
		totalPage = 1
	} else {
		totalPage = totalPage1
	}

	nextPageString := fmt.Sprintf("%s/%s/%spage=%d&limit=%d", conf.App.Url, conf.App.Name_api, endpoint, nextPage, limit)
	previousPageUrlString := fmt.Sprintf("%s/%s/%spage=%d&limit=%d", conf.App.Url, conf.App.Name_api, endpoint, previousPage, limit)
	firstPageUrlString := fmt.Sprintf("%s/%s/%spage=%d&limit=%d", conf.App.Url, conf.App.Name_api, endpoint, 1, limit)
	lastPageUrlString := fmt.Sprintf("%s/%s/%spage=%v&limit=%d", conf.App.Url, conf.App.Name_api, endpoint, totalPage, limit)

	results := helperDatabases.ResponseBackPaginationDTO{
		TotalData:        &totalData,
		TotalDataPerPage: &limit,
		CurrentPage:      page,
		PreviousPage:     &previousPage,
		TotalPage:        &totalPage,
		NextPageUrl:      &nextPageString,
		PreviousPageUrl:  &previousPageUrlString,
		FirstPageUrl:     &firstPageUrlString,
		LastPageUrl:      &lastPageUrlString,
	}

	return results

}

func DriverIncentive(totalCost int) float32 {
	driverIncentive := (0.05 * float32(totalCost))
	return driverIncentive
}

func TotalCost(days uint, carDailyCost uint) uint {
	totalCost := (days * carDailyCost)
	return totalCost
}

func Days(endTime time.Time, startTime time.Time) uint {
	days := uint(endTime.Sub(startTime).Hours() / 24)
	return days
}

func Discount(totalCost float32, discountValue float32) float32 {
	discountTemp := (totalCost * discountValue)
	return discountTemp
}

func TotalDriverCost(days uint, dailyCostDriver uint) uint {
	totalDriverCostTemp := (days * dailyCostDriver)
	return totalDriverCostTemp
}

func CompareString(str1, str2 string) bool {
	return str1 == str2
}
