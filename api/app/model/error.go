package model

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type ErrorMessageST struct {
	Message    string        `json:"error" validate:"required"`
	Parameters []interface{} `json:"parameters" validate:"required"`
} // @name ErrorMessage

func NewErrorMessage(message string, parameters ...interface{}) *ErrorMessageST {
	if parameters == nil {
		parameters = make([]interface{}, 0)
	}
	return &ErrorMessageST{
		Message:    message,
		Parameters: parameters,
	}
}

type ErrorST struct {
	StatusCode int                          `json:"-"`
	Errors     map[string][]*ErrorMessageST `json:"errors" validate:"required"`
} // @name Errors

func NewError(statusCode int) *ErrorST {
	return &ErrorST{
		StatusCode: statusCode,
		Errors:     make(map[string][]*ErrorMessageST),
	}
}

func (e *ErrorST) AddError(name string, message string, parameters ...interface{}) *ErrorST {
	if e.Errors[name] == nil {
		e.Errors[name] = make([]*ErrorMessageST, 0)
	}
	e.Errors[name] = append(e.Errors[name], NewErrorMessage(message, parameters...))
	return e
}

func (e *ErrorST) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *ErrorST) Send(c *fiber.Ctx) error {
	return c.Status(e.StatusCode).JSON(e)
}

func (e *ErrorST) Error() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
