package errors

import (
	"fmt"
	"log"

	"github.com/ztrue/tracerr"
)

type InternalError struct {
	Message string
	Args    []any
}

func (err InternalError) Error() string {
	return fmt.Sprintf(err.Message, err.Args...)
}

func NewTrace(message string, args ...any) error {
	return WrapTrace(InternalError{Message: message, Args: args})
}

func New(message string, args ...any) error {
	return InternalError{Message: message, Args: args}
}

func UnwrapTrace(err error) error {
	return tracerr.Unwrap(err)
}

func WrapTrace(err error) error {
	return tracerr.Wrap(err)
}

func SprintTrace(err error) string {
	return tracerr.Sprint(err)
}

func LogTraceIfErr(err error) {
	if err != nil {
		log.Println(SprintTrace(WrapTrace(err)))
	}
}
