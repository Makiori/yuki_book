package controller

type PaginationQueryBody struct {
	Page     uint `json:"page" form:"page"`
	PageSize uint `json:"pageSize" form:"pageSize"`
}
