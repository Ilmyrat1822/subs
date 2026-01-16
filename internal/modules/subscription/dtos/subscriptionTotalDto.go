package dtos

type TotalCostResponse struct {
	Total int64
	Count int64
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
}
