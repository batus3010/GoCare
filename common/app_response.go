package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
}

// NewSuccessResponse returns detailed response
func NewSuccessResponse(data, paging interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging}
}

// SimpleSuccessResponse is a simple response with just data
func SimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{Data: data}
}
