package asyncjob

import (
	"context"
	"time"
)

// Job Requirement:
// 1. Job can do something (handler)
// 2. Job can retry
//  2.1 Config retry times and duration
// 3. Should be stateful (status)
// 4. We should have job manager to manage jobs (*)

type Job interface {
	Execute(ctx context.Context) error // thực thi
	Retry(ctx context.Context) error   // thử lại
	State() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout = time.Second * 10
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
)

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

func (js JobState) String() string {
	// gán list const trên có value js như phía dưới ~~ lấy mảng truy xuất index
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		retryIndex: -1,
		state:      StateInit,
		stopChan:   make(chan bool),
	}

	return &j
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning // running

	var err error
	err = j.handler(ctx)

	if err != nil {
		j.state = StateFailed // fail
		return err
	}

	j.state = StateCompleted // completed

	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1 // tăng retryIndex
	time.Sleep(j.config.Retries[j.retryIndex])

	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) RetryIndex() int {
	return j.retryIndex
}

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}

	j.config.Retries = times
}
