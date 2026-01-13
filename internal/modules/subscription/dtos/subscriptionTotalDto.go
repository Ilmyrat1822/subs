package dtos

type TotalCostResponse struct {
	TotalCost int    `json:"total_cost"`
	Period    string `json:"period"`
	Count     int    `json:"count"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
