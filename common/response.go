package common

type successRes struct {
	StatusCode *int        `json:"status_code"`
	Message    *string     `json:"message"`
	Data       interface{} `json:"data"`
	Paging     interface{} `json:"paging,omitempty"`
	Filter     interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}, statusCode int, message string) *successRes {

	return &successRes{
		Data:       data,
		Paging:     paging,
		Filter:     filter,
		StatusCode: &statusCode,
		Message:    &message,
	}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil, 200, "Successful!")
}
