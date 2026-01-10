package errors

import (
	"errors"
	"net/http"

	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
)

var ErrGenericInternalServerErrorKey = "error:player:internal_server_error"
var ErrUserNotFound = "error:player:not_found"
var ErrBadRequestKey = "error:player:bad_request"

var (
	ErrorGenericInternalServerError = errors.New(ErrGenericInternalServerErrorKey)
	ErrorBadRequest                 = errors.New(ErrBadRequestKey)
	ErrorUserNotFound               = errors.New(ErrUserNotFound)
)

var ErrorDetailsMap = map[string]errorhandler.ErrorDetails{
	ErrGenericInternalServerErrorKey: errorhandler.ErrorDetails{"Something went wrong, please try again later.", http.StatusInternalServerError},
	ErrBadRequestKey:                 errorhandler.ErrorDetails{"Bad request. Please check your input.", http.StatusBadRequest},
	ErrUserNotFound:                  errorhandler.ErrorDetails{"Player not found.", http.StatusNotFound},
}
