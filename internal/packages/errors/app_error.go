package errors

import "net/http"

type AppError struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Detail  interface{} `json:"detail"`
}

func (e *AppError) Error() string {
    return e.Message
}

func NewBadRequest(msg string, detail interface{}) *AppError {
    return &AppError{
        Code:    http.StatusBadRequest,
        Message: msg,
        Detail:  detail,
    }
}

func NewUnauthorized(msg string) *AppError {
    return &AppError{
        Code:    http.StatusUnauthorized,
        Message: msg,
        Detail:  nil,
    }
}

func NewForbidden(msg string) *AppError {
    return &AppError{
        Code:    http.StatusForbidden,
        Message: msg,
        Detail:  nil,
    }
}

func NewNotFound(msg string) *AppError {
    return &AppError{
        Code:    http.StatusNotFound,
        Message: msg,
        Detail:  nil,
    }
}

func NewInternal(detail interface{}) *AppError {
    return &AppError{
        Code:    http.StatusInternalServerError,
        Message: "internal server error",
        Detail:  detail,
    }
}
