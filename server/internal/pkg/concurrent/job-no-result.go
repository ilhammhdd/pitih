package concurrent

type JobNoResult interface {
	Do(work func() error)
	Job
}

func NewJobNoResult() JobNoResult {
	return &jobNoResult{job: newJob()}
}

type jobNoResult struct {
	*job
}

func (jnr *jobNoResult) Do(work func() error) {
	jnr.wg.Add(1)
	GoRecover(func() {
		defer jnr.wg.Done()
		if err := work(); err != nil {
			jnr.errs = append(jnr.errs, err)
		}
	})
}
