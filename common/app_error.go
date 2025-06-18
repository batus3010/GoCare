package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

// NewErrorResponse is used for most errors, just return Bad Request
func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorizedErrorResponse(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

// NewCustomErrorResponse is used for errors without root
func NewCustomErrorResponse(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

// RootError traces back to root error
func (e *AppError) RootError() error {
	var err *AppError
	// if error can be cast as AppError then return parent Error
	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}
	return e.RootErr
}

// Every struct implements Error() is called error
func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrorDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with db", err.Error(), "ErrDB")
}

func ErrorInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrorInvalidRequest(err error) *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, err, "invalid request", err.Error(), "INVALID_REQUEST")
}

func ErrorCannotListEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotList%s", entity))
}

func ErrorCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotDelete%s", entity))
}

func ErrorCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotUpdate%s", entity))
}

func ErrorCannotGetEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotGet%s", entity))
}

func ErrorEntityExisted(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Error%sExisted", entity))
}

func ErrorEntityNotFound(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Error%sNotFound", entity))
}

func ErrorCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotCreate%s", entity))
}

func ErrorNoPermission(err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("you have no permission to do this"),
		fmt.Sprintf("ErrorNoPermission"))
}
