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

func HandleNoAllowMethod(c *gin.Context) {
	handleErr := NoAllowMethod()
	c.JSON(handleErr.Code, handleErr)
	return
}

const (
	INTERNA_ERROR   = 1000
	NOT_FOUND       = 1001
	NO_ALLOW_METHDO = 1002
	UNKNOWN_ERROR   = 1003
	PARAMETER_ERROR = 1004
)

type HttpException struct {
	Code       int    `json:"-"`
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
}

func (e *HttpException) Error() string {
	return e.Msg
}

type HandlerFunc func(c *gin.Context) error

func ErrHandler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		err = h(c)
		if err != nil {
			var apiException *HttpException
			if h, ok := err.(*HttpException); ok {
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

func newHttpException(code int, errorCode int, msg string) *HttpException {
	return &HttpException{
		Code:       code,
		StatusCode: errorCode,
		Msg:        msg,
	}
}

// InternalError
// Service Internal Error response
func InternalError() *HttpException {
	return newHttpException(http.StatusInternalServerError, INTERNA_ERROR, http.StatusText(http.StatusInternalServerError))
}

// NotFound
// Not found page error response
func NotFound() *HttpException {
	return newHttpException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
}

// NoAllowMethod
// Not Allow Method
func NoAllowMethod() *HttpException {
	return newHttpException(http.StatusMethodNotAllowed, NO_ALLOW_METHDO, http.StatusText(http.StatusMethodNotAllowed))
}

// UnknownError
// Unknown Error response
func UnknownError(message string) *HttpException {
	return newHttpException(http.StatusForbidden, UNKNOWN_ERROR, message)
}
