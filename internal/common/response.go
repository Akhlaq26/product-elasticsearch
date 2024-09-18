package common

// @description Represents a BaseResponse
type BaseResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewBaseResponse[T any](message string, data T) *BaseResponse[T] {
	return &BaseResponse[T]{
		Message: message,
		Data:    data,
	}
}
