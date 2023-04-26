package helperDatabases

type ResponseBackPaginationDTO struct {
	TotalData        *int     `json:"total_data"`
	TotalDataPerPage *int     `json:"total_data_per_page"`
	CurrentPage      *int     `json:"current_page"`
	PreviousPage     *int     `json:"previous_page"`
	TotalPage        *float64 `json:"total_page"`
	NextPageUrl      *string  `json:"next_page_url"`
	PreviousPageUrl  *string  `json:"previous_page_url"`
	FirstPageUrl     *string  `json:"first_page_url"`
	LastPageUrl      *string  `json:"last_page_url"`
}
