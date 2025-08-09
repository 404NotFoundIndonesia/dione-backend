package dto

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateErrorResponse(message string) Response[any] {
	return Response[any]{
		Message: message,
		Data:    nil,
	}
}

func CreateErrorResponseWithPayload(message string, data any) Response[any] {
	return Response[any]{
		Message: message,
		Data:    data,
	}
}

func CreateSuccessResponse[T any](data T) Response[T] {
	return Response[T]{
		Message: "success",
		Data:    data,
	}
}
