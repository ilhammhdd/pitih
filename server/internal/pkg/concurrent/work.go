package concurrent

import (
	"fmt"

	"ilhammhdd.com/pitih/internal/pkg/errors"
)

func GoRecover(work func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errors.LogTraceIfErr(errors.New(fmt.Sprintf("%+v", r)))
			}
		}()
		work()
	}()
}

func GoRecoverWithParam[T any](work func(T), param T) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errors.LogTraceIfErr(errors.New(fmt.Sprintf("%+v", r)))
			}
		}()
		work(param)
	}()
}

// this function will receive blocking the resultChan
func Receive[T any](resultChan <-chan any) T {
	var finalResult T
	result := <-resultChan
	if result != nil {
		assertedResult, ok := result.(T)
		if ok {
			finalResult = assertedResult
		}
	}
	return finalResult
}
