package errors

import (
	"errors"
	"net/http"

	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
)

var ErrGenericInternalServerErrorKey = "error:user:internal_server_error"
var ErrUserNotFound = "error:user:not_found"
var ErrBadRequestKey = "error:user:bad_request"
var ErrUserPlayerMappingNotFoundKey = "error:user:player_mapping_not_found"
var ErrCannotFollowSelfKey = "error:user:cannot_follow_self"

var (
	ErrorGenericInternalServerError = errors.New(ErrGenericInternalServerErrorKey)
	ErrorBadRequest                 = errors.New(ErrBadRequestKey)
	ErrorUserNotFound               = errors.New(ErrUserNotFound)
	ErrorUserPlayerMappingNotFound  = errors.New(ErrUserPlayerMappingNotFoundKey)
	ErrorCannotFollowSelf           = errors.New(ErrCannotFollowSelfKey)
)

var ErrorDetailsMap = map[string]errorhandler.ErrorDetails{
	ErrGenericInternalServerErrorKey: {Message: "Something went wrong, please try again later.", Code: http.StatusInternalServerError},
	ErrBadRequestKey:                 {Message: "Bad request. Please check your input.", Code: http.StatusBadRequest},
	ErrUserNotFound:                  {Message: "User not found.", Code: http.StatusNotFound},
	ErrUserPlayerMappingNotFoundKey:  {Message: "Please choose the player.", Code: http.StatusNotFound},
	ErrCannotFollowSelfKey:           {Message: "You cannot follow yourself.", Code: http.StatusBadRequest},
	ErrGenericInternalServerErrorKey: errorhandler.ErrorDetails{"Something went wrong, please try again later.", http.StatusInternalServerError},
	ErrBadRequestKey:                 errorhandler.ErrorDetails{"Bad request. Please check your input.", http.StatusBadRequest},
	ErrUserNotFound:                  errorhandler.ErrorDetails{"User not found.", http.StatusNotFound},
}
