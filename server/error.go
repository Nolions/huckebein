package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	c.JSON(handleErr.Code, handleErr)
	return
}

const (
	INTERNA_ERROR = 1000
	NOT_FOUND     = 1001
	UNKNOWN_ERROR   = 1002
	PARAMETER_ERROR = 1003
)

type APIException struct {
	Code       int    `json:"-"`
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
}

func (e *APIException) Error() string {
	return e.Msg
}

type HandlerFunc func(c *gin.Context) error

func ErrHandler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		err = h(c)
		if err != nil {
			var apiException *APIException
			if h, ok := err.(*APIException); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				apiException = UnknownError(e.Error())
			} else {
				apiException = InternalError()
			}
			c.JSON(apiException.Code, apiException)
			return
		}
	}
}

func newAPIException(code int, errorCode int, msg string) *APIException {
	return &APIException{
		Code:       code,
		StatusCode: errorCode,
		Msg:        msg,
	}
}

// InternalError
// Service Internal Error response
func InternalError() *APIException {
	return newAPIException(http.StatusInternalServerError, INTERNA_ERROR, http.StatusText(http.StatusInternalServerError))
}

// NotFound
// Not found page error response
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
}

// UnknownError
// Unknown Error response
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, UNKNOWN_ERROR, message)
}
