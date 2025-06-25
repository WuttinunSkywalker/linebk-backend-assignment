package response

import "math"

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewSuccess(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Data:    data,
	}
}

func NewError(message string) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Message: message,
	}
}

func NewPaginated(items interface{}, totalItems int, page, limit int) *PaginatedResponse {
	return &PaginatedResponse{
		Success: true,
		Data:    items,
		Pagination: &Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: int(math.Ceil(float64(totalItems) / float64(limit))),
		},
	}
}
