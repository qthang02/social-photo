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

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func NewFullErrorResponse(statusCode int, root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, message, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    message,
		Key:        key,
	}
}

func NewCustomError(root error, message, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, message, root.Error(), key)
	}

	return NewErrorResponse(errors.New(message), message, message, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "INTERNAL_ERROR")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_LIST_%s", strings.ToUpper(entity)))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_DELETE_%s", strings.ToUpper(entity)))
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_UPDATE_%s", strings.ToUpper(entity)))
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_CREATE_%s", strings.ToUpper(entity)))
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_GET_%s", strings.ToUpper(entity)))
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("%s has been deleted", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_DELETED", strings.ToUpper(entity)))
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("%s has been existed", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_EXISTED", strings.ToUpper(entity)))
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_NOT_FOUND", strings.ToUpper(entity)))
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err,
		"you don't have permission to do this action",
		"ERR_NO_PERMISSION")
}

var RecordNotFound = errors.New("record not found")
