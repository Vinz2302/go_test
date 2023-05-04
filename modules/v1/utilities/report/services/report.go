package services

import (
	"fmt"
	model "rest-api/modules/v1/utilities/report/models"
	repo "rest-api/modules/v1/utilities/report/repository"
)

type IReportService interface {
	FindMonthlyCompanyIncome(year int, month int) ([]model.Report, error)
	FindBookingActivity(year int, month int) ([]model.ReportCar, error)
	FindDriverActivity(year int, month int) ([]model.ReportDriver, error)
}

type reportService struct {
	repository repo.IReportRepository
}

func NewReportService(repository repo.IReportRepository) *reportService {
	return &reportService{repository}
}

func (s *reportService) FindMonthlyCompanyIncome(year int, month int) ([]model.Report, error) {
	report, errReport := s.repository.FindMonthlyCompanyIncome(year, month)
	if errReport != nil {
		return nil, errReport
	}
	fmt.Println("report 2 = ", *&report)

	/* totalDriverExpense := uint(report.TotalDriverCost + report.TotalDriverIncentive)
	report.TotalDriverExpense = &totalDriverExpense

	totalGrossIncome := uint(report.TotalCost - uint(report.Discount) + report.TotalDriverCost)
	report.TotalGrossIncome = &totalGrossIncome

	totalNettIncome := uint(report.TotalCost - uint(report.Discount) + report.TotalDriverCost - report.TotalDriverIncentive)
	report.TotalNettIncome = &totalNettIncome */

	for i := range report {
		fmt.Println(*report[i].Discount)
		totalDriverExpense := *report[i].TotalDriverCost + *report[i].TotalDriverIncentive
		report[i].TotalDriverExpense = &totalDriverExpense

		totalGrossIncome := *report[i].TotalCost - uint(*report[i].Discount) + *report[i].TotalDriverCost
		report[i].TotalGrossIncome = &totalGrossIncome

		totalNettIncome := *report[i].TotalCost - uint(*report[i].Discount) + *report[i].TotalDriverCost - *report[i].TotalDriverIncentive
		report[i].TotalNettIncome = &totalNettIncome
	}

	return report, nil
}

func (s *reportService) FindBookingActivity(year int, month int) ([]model.ReportCar, error) {
	reportCar, errReport := s.repository.FindBookingActivity(year, month)
	if errReport != nil {
		return nil, errReport
	}
	return reportCar, nil
}

func (s *reportService) FindDriverActivity(year int, month int) ([]model.ReportDriver, error) {
	reportDriver, errReport := s.repository.FindDriverActivity(year, month)
	if errReport != nil {
		return nil, errReport
	}
	return reportDriver, nil
}
