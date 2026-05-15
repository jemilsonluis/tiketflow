package dto

type ResponseDTO struct {
	Data       interface{}
	IsSuccess  bool
	StatusCode int
	Message    string
}

func NewResponseDTO(data interface{}, msg string, statusCode int, isSuccess bool) *ResponseDTO {
	return &ResponseDTO{
		Data:       data,
		IsSuccess:  isSuccess,
		StatusCode: statusCode,
		Message:    msg,
	}
}
