package models

// PaginationResponse - Standard pagination response
// @name PaginationResponse
type PaginationResponse struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
	Total int `json:"total" example:"100"`
	Pages int `json:"pages" example:"10"`
}

// ErrorResponse - Standard error response
// @name ErrorResponse
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// SuccessResponse - Standard success response
// @name SuccessResponse
type SuccessResponse struct {
	Status string `json:"status" example:"success"`
}

// ListResponse - Standard list response with pagination
// @name ListResponse
type ListResponse struct {
	Status     string             `json:"status" example:"success"`
	Datas      interface{}        `json:"datas"`
	Pagination PaginationResponse `json:"pagination"`
}

type ListDashboardResponse struct {
	Datas      interface{}        `json:"datas"`
	Pagination PaginationResponse `json:"pagination"`
}

type ListDetailResponse struct {
	Datas   interface{} `json:"datas"`
	Message string      `json:"messagae"`
}
