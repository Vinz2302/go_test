package helperDatabases

type QueryParamPaginationEntity struct {
	Page    *int    `form:"page"`
	Limit   *int    `form:"limit"`
	Offset  *int    `form:"offet"`
	Search  *string `form:"search"`
	OrderBy *string `form:"order_by`
}
