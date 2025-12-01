// By Ratheesh Kumar Golang Developer
package errors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Standardized error messages
const (
	BadRequest                 = "bad request"
	InternalServerError        = "internal server error"
	AlreadyExists              = "already exists"
	NotFound                   = "not Found"
	Unauthorized               = "unauthorized"
	Unauthenticated            = "unauthenticated"
	Forbidden                  = "forbidden"
	BadQueryParams             = "invalid query params"
	TooManyRequests            = "too many request"
	InvalidField               = "invalid field"
	RequiredField              = "required field"
	RequiredParams             = "param is required"
	RequestTimeout             = "request timeout"
	MissingAuthoirzationHeader = "missing Authorization header"
	BasicAuth                  = "authentication error: please provide valid basic authentication credentials in the 'Authorization' header"
	BearerToken                = "authentication error: please provide a valid bearer token in the 'Authorization' header"
	InvalidToken               = "expired or invalid token"
	InvalidSession             = "expired or invalid session"
	SessionUsed                = "session is already used"
	EndpointNotFound           = "the endpoint you requested doesn't exist on server"
)

// HTTPErrorResponse represents a standardized error response structure
var (
	ErrBadRequest          = errors.New(BadRequest)
	ErrInternalServerError = errors.New(InternalServerError)
	ErrAlreadyExists       = errors.New(AlreadyExists)
	ErrNotFound            = errors.New(NotFound)
	ErrUnauthorized        = errors.New(Unauthorized)
	ErrUnauthenticated     = errors.New(Unauthenticated)
	ErrForbidden           = errors.New(Forbidden)
	ErrBadQueryParams      = errors.New(BadQueryParams)
	ErrTooManyRequests     = errors.New(TooManyRequests)
	ErrInvalidField        = errors.New(InvalidField)
	ErrRequiredField       = errors.New(RequiredField)
	ErrRequiredParams      = errors.New(RequiredParams)
	ErrRequestTimeout      = errors.New(RequestTimeout)
	ErrMissingAuthHeader   = errors.New(MissingAuthoirzationHeader)
	ErrBasicAuth           = errors.New(BasicAuth)
	ErrBearerToken         = errors.New(BearerToken)
	ErrInvalidToken        = errors.New(InvalidToken)
	ErrInvalidSession      = errors.New(InvalidSession)
	ErrSessionUsed         = errors.New(SessionUsed)
	ErrEndpointNotFound    = errors.New(EndpointNotFound)
)

// HTTPErrorResponse represents a standardized error response structure
type HTTPErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

// BadqueryParamsError creates a bad query params error response
func BadqueryParamsError(message string) *HTTPErrorResponse {
	return &HTTPErrorResponse{
		StatusCode: fiber.StatusBadRequest,
		Error:      BadQueryParams,
		Message:    message,
	}
}

// InternalServerErrorResponse creates an internal server error response
func InternalServerErrorResponse(message string) *HTTPErrorResponse {
	return &HTTPErrorResponse{
		StatusCode: fiber.StatusInternalServerError,
		Error:      InternalServerError,
		Message:    message,
	}
}

// NotFoundErrorResponse creates a not found error response
func NotFoundErrorResponse(message string) *HTTPErrorResponse {
	return &HTTPErrorResponse{
		StatusCode: fiber.StatusNotFound,
		Error:      NotFound,
		Message:    message,
	}
}

// UnauthorizedErrorResponse creates an unauthorized error response
func UnauthorizedErrorResponse(message string) *HTTPErrorResponse {
	return &HTTPErrorResponse{
		StatusCode: fiber.StatusUnauthorized,
		Error:      Unauthorized,
		Message:    message,
	}
}

// BadRequestErrorResponse creates a bad request error response
func BadRequestErrorResponse(message string) *HTTPErrorResponse {
	return &HTTPErrorResponse{
		StatusCode: fiber.StatusBadRequest,
		Error:      BadRequest,
		Message:    message,
	}
}

// NewDoesntExistError creates a new "doesn't exist" error for a given entity
func NewDoesntExistError(entity string) error {
	return errors.New(entity + " doesn't exist")
}
