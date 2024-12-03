package concurrent

// Don't forget to call Close, prefer deferred call
type JobWithResult interface {
	DoWithResult(work func() (any, error)) <-chan any
	Job
}

func NewJobWithResult() JobWithResult {
	return &jobWithResult{job: newJob()}
}

type jobWithResult struct {
	resultChans []chan any
	*job
}

func (jwr *jobWithResult) DoWithResult(work func() (any, error)) <-chan any {
	jwr.wg.Add(1)
	resultChan := make(chan any)
	jwr.resultChans = append(jwr.resultChans, resultChan)
	GoRecover(func() {
		result, err := work()
		if err != nil {
			jwr.errs = append(jwr.errs, err)
		}
		GoRecover(func() {
			resultChan <- result
		})
		jwr.wg.Done()
	})
	return resultChan
}
