package concurrent

import (
	"strings"
	"sync"

	"ilhammhdd.com/pitih/internal/pkg/errors"
)

type Job interface {
	WaitErrors() []error
	Wait() error
	errorsJoined() error
}

func NewJob() Job {
	return newJob()
}

func newJob() *job {
	return &job{wg: &sync.WaitGroup{}}
}

type job struct {
	wg   *sync.WaitGroup
	errs []error
}

func (j *job) Wait() error {
	j.wg.Wait()
	return j.errorsJoined()
}

func (j *job) WaitErrors() []error {
	j.wg.Wait()
	return j.errs
}

func (j *job) errorsJoined() error {
	if len(j.errs) == 0 {
		return nil
	}
	var errMessages []string
	for idx := range j.errs {
		errMessages = append(errMessages, j.errs[idx].Error())
	}
	return errors.NewTrace(strings.Join(errMessages, " | "))
}
