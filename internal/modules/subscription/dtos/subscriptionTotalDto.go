package dtos

type TotalCostResponse struct {
	TotalCost int    `json:"total_cost" example:"1200"`
	Period    string `json:"period" example:"01-2025 to 12-2025"`
	Count     int    `json:"count" example:"3"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
}
