package response

type Response[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewResponse[T any](succcess bool, data T, error string) Response[T] {
	return Response[T]{
		Success: succcess,
		Data:    data,
		Error:   error,
	}
}
