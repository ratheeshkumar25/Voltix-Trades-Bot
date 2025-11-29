// By Ratheesh Kumar Golang Developer
package http

import (
	"encoding/json"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"
	"ratheeshkumar25/github.com/trading_bot/auth-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

// Locals constants
const (
	LocalsAllowed = "allowed"
	LocalsClient  = "client"
	LocalsToken   = "token"
	LocalsDevice  = "device"
	LocalsOs      = "os"
	LocalsChannel = "channel"
)

const (
	StatusBadRequest          = fiber.StatusBadRequest
	StatusUnauthorized        = fiber.StatusUnauthorized
	StatusForbidden           = fiber.StatusForbidden
	StatusNotFound            = fiber.StatusNotFound
	StatusInternalServerError = fiber.StatusInternalServerError
	StatusOK                  = fiber.StatusOK
	StatusCreated             = fiber.StatusCreated
	StatusNoContent           = fiber.StatusNoContent
)

const (
	ErrBadRequest          = "Bad request"
	ErrInternalServerError = "Internal server error"
	ErrAlreadyExists       = "Already exists"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrForbidden           = "Forbidden"
	ErrBadQueryParams      = "Invalid query params"
	ErrRequestTimeout      = "Request Timeout"
	ErrEndpointNotFound    = "The endpoint you requested doesn't exist on server"
)

type App struct {
	//fiber app instance
	Fiber *fiber.App
	// Logger instance
	Log *logger.Logger
}

// NewApp creates a new App instance with Fiber and Logger
func NewApp(fiberApp *fiber.App, log *logger.Logger) *App {
	newApp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	return &App{
		Fiber: newApp,
		Log:   log,
	}
}

// HttpResponse represents a standardized HTTP response structure
type HttpResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type WSResponse struct {
	EventReason string
	Success     bool        `json:"success"`
	Data        interface{} `json:"data,omitempty"`
	Error       string      `json:"error,omitempty"`
	Message     string      `json:"message,omitempty"`
}

// HTTP 200 OK response
func (a *App) HttpResponseok(c *fiber.Ctx, data interface{}) *HttpResponse {
	return &HttpResponse{
		Success: true,
		Code:    StatusOK,
		Data:    data,
		Error:   "",
		Message: "",
	}
}

// HTTP 201 Created response
func (a *App) HttpResponseCreated(c *fiber.Ctx, data interface{}) *HttpResponse {
	return &HttpResponse{
		Success: true,
		Code:    StatusCreated,
		Data:    data,
		Error:   "",
		Message: "",
	}
}

// HTTP 204 No Content response
// http 204 no content http response
func (a *App) HttpResponseNoContent(c *fiber.Ctx) error {
	return c.Status(StatusNoContent).JSON(
		&HttpResponse{
			Success: true,
			Code:    StatusNoContent,
			Data:    nil,
			Error:   "",
			Message: "",
		})
}

// HTTP 400 Bad Request response

// http 400 bad request http response
func (a *App) HttpResponseBadRequest(c *fiber.Ctx, message error) error {
	a.Log.Logger.Error(message.Error())
	return c.Status(StatusBadRequest).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusBadRequest,
			Data:    nil,
			Error:   ErrBadRequest,
			Message: message.Error(),
		})
}

// http 400 bad query params http response
func (a *App) HttpResponseBadQueryParams(c *fiber.Ctx, message error) error {
	a.Log.Logger.Error(message.Error())
	return c.Status(StatusBadRequest).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusBadRequest,
			Data:    nil,
			Error:   ErrBadQueryParams,
			Message: message.Error(),
		})
}

// http 404 not found http response
func (a *App) HttpResponseNotFound(c *fiber.Ctx, message error) error {
	a.Log.Logger.Error(message.Error())
	return c.Status(StatusNotFound).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusNotFound,
			Data:    nil,
			Error:   ErrNotFound,
			Message: message.Error(),
		})
}

// http 500 internal server error response
func (a *App) HttpResponseInternalServerErrorRequest(c *fiber.Ctx, message error) error {
	a.Log.Logger.Error(message.Error())
	return c.Status(StatusInternalServerError).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusInternalServerError,
			Data:    nil,
			Error:   ErrInternalServerError,
			Message: message.Error(),
		})
}

// http 403 The client does not have access rights to the content;
// that is, it is unauthorized, so the server is refusing to give the requested resource
func (a *App) HttpResponseForbidden(c *fiber.Ctx, message error) error {
	a.Log.Logger.Error(message.Error())
	return c.Status(StatusForbidden).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusForbidden,
			Data:    nil,
			Error:   ErrForbidden,
			Message: message.Error(),
		})
}

// http 401 the client must authenticate itself to get the requested response
func (a *App) HttpResponseUnauthorized(c *fiber.Ctx, message error) error {
	a.Log.Logger.Error(message.Error())
	return c.Status(StatusUnauthorized).JSON(
		&HttpResponse{
			Success: false,
			Code:    StatusUnauthorized,
			Data:    nil,
			Error:   ErrUnauthorized,
			Message: message.Error(),
		})
}

// http 200 retrieve File response
func (a *App) HttpResponseFile(c *fiber.Ctx, file []byte) error {
	return c.Status(fiber.StatusOK).Send(file)
}

// WS 200 ok http response
func (a *App) WSResponseOK(event models.EventType, data interface{}) *models.Event {
	return &models.Event{
		Type:    event,
		Payload: data,
	}
}

// WS 400 bad request http response
func (a *App) WSResponseBadRequest(event models.EventType, message error) *models.Event {
	a.Log.Logger.Error(message.Error())
	return &models.Event{
		Type: models.EventBadRequest,
		Payload: models.ErrorPayload{
			Message: message.Error(),
			Reason:  event,
		},
	}
}

// WS 500 internal server error response
func (a *App) WSResponseInternalServerErrorRequest(event models.EventType, message error) *models.Event {
	a.Log.Logger.Error(message.Error())
	return &models.Event{
		Type: models.EventInternalServerError,
		Payload: models.ErrorPayload{
			Message: message.Error(),
			Reason:  event,
		},
	}
}

// WS 404 bad request http response
func (a *App) WSResponseNotFound(event models.EventType, message error) *models.Event {
	a.Log.Logger.Error(message.Error())
	return &models.Event{
		Type: models.EventNotFound,
		Payload: models.ErrorPayload{
			Message: message.Error(),
			Reason:  event,
		},
	}
}

// WS 403 The client does not have access rights to the content;
// that is, it is unauthorized, so the server is refusing to give the requested resource
func (a *App) WSResponseForbidden(event models.EventType, message error) *models.Event {
	a.Log.Logger.Error(message.Error())
	return &models.Event{
		Type: models.EventForbidden,
		Payload: models.ErrorPayload{
			Message: message.Error(),
			Reason:  event,
		},
	}
}

// ws 401 the client must authenticate itself to get the requested response
func (a *App) WSResponseUnauthorized(event models.EventType, message error) *models.Event {
	a.Log.Logger.Error(message.Error())
	return &models.Event{
		Type: models.EventUnauthorized,
		Payload: models.ErrorPayload{
			Message: message.Error(),
			Reason:  event,
		},
	}
}
