package eaapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessageError() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{Code: http.StatusForbidden, Message: message}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: message}
}

func NewAlreadyExistError(message string) *AppError {
	return &AppError{Code: http.StatusConflict, Message: message}
}

func NewValidationError(message string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: message}
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FailedResponse struct {
	Status  ResponseStatus `json:"status"`
	Code    int            `json:"code"`
	Message string         `json:"message"`
}

func ResponseHandler(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.AbortWithStatusJSON(code, SuccessResponse{
		Status:  ResponseStatusSuccess,
		Message: message,
		Data:    data,
	})
}

func ErrorResponseHandler(ctx *gin.Context, err *AppError) {
	ctx.AbortWithStatusJSON(err.Code, FailedResponse{
		Status:  ResponseStatusFailed,
		Code:    err.Code,
		Message: err.Message,
	})
}

func NewHandlerNotFoundError(ctx *gin.Context, err *AppError) {
	if err == nil {
		err = NewNotFoundError(ErrNotFound)
	}

	ErrorResponseHandler(ctx, err)
}

func NewHandlerUnexpectedError(ctx *gin.Context, err *AppError) {
	if err == nil {
		err = NewUnexpectedError(ErrInternalServerError)
	}

	ErrorResponseHandler(ctx, err)
}

func NewHandlerForbiddenError(ctx *gin.Context, err *AppError) {
	if err == nil {
		err = NewForbiddenError(ErrForbiddenAccess)
	}

	ErrorResponseHandler(ctx, err)
}

func NewHandlerBadRequestError(ctx *gin.Context, err *AppError) {
	if err == nil {
		err = NewBadRequestError(ErrBadRequest)
	}

	ErrorResponseHandler(ctx, err)
}

func NewHandlerAlreadyExistError(ctx *gin.Context, err *AppError) {
	if err == nil {
		err = NewAlreadyExistError(ErrConflict)
	}

	ErrorResponseHandler(ctx, err)
}

func NewHandlerValidationError(ctx *gin.Context, err *AppError) {
	if err == nil {
		err = NewValidationError(ErrBadRequest)
	}

	ErrorResponseHandler(ctx, err)
}

func NewHandlerError(ctx *gin.Context, err *AppError) {
	if err == nil {
		err = NewValidationError(ErrBadRequest)
	}

	ErrorResponseHandler(ctx, err)
}
