package driver

import (
	handler "rest-api/modules/v1/utilities/report/handler"
	repo "rest-api/modules/v1/utilities/report/repository"
	service "rest-api/modules/v1/utilities/report/services"
)

var (
	ReportRepository = repo.NewReportRepository(DB)
	ReportService    = service.NewReportService(ReportRepository)
	ReportHandler    = handler.NewReportHandler(ReportService)
)
