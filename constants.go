package eaapi

type HandlerError = string

const (
	ErrBadRequest          HandlerError = "bad request"
	ErrNotFound            HandlerError = "not found"
	ErrForbiddenAccess     HandlerError = "forbidden access"
	ErrConflict            HandlerError = "conflict"
	ErrInternalServerError HandlerError = "internal server error"
)

type ResponseStatus = string

const (
	ResponseStatusSuccess ResponseStatus = "success"
	ResponseStatusFailed  ResponseStatus = "failed"
)
