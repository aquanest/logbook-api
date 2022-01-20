package usecase

import (
	"logbook-api/internal/config"
	"logbook-api/internal/infrastructure"
	"logbook-api/pkg/http"

	"github.com/labstack/echo/v4"
)

// Usecase struct
type Usecase struct {
	Config     *config.Config
	Repository *infrastructure.Repository
}

// GetHealth returns the message
func (u *Usecase) GetHealth() (http.BaseResponse, error) {
	return http.BaseResponse{
		Code:    http.StatusOK,
		Message: http.MessageOK,
	}, nil
}

// NotFound returns the message
func (u *Usecase) NotFound() (http.BaseResponse, error) {
	return http.BaseResponse{
		Code:    http.StatusNotFound,
		Message: http.MessageNotFound,
	}, nil
}

// MethodNotAllowed returns the message
func (u *Usecase) MethodNotAllowed() (http.BaseResponse, error) {
	return http.BaseResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: http.MessageMethodNotAllowed,
	}, nil
}

// UnexpectedError returns the message
func (u *Usecase) UnexpectedError(err *echo.HTTPError) (http.BaseResponse, error) {
	return http.BaseResponse{
		Code:    http.StatusInternalServerError,
		Message: http.MessageInternalServerError + ". " + err.Message.(string),
	}, nil
}
