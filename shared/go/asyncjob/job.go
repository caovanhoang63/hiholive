package asyncjob

import (
	"context"
	"time"
)

// Job Requirement:
// 1. Job can do something (handler)
// 2. Job can retry
// 	2.1 Config retry times and duration
// 3. Should be stateful
// 4. We should have job manager to manage jobs (*)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout    = time.Second * 10
	defaultMaxRetryCount = 3
)

var (
	defaultRetryTime = []time.Duration{
		time.Second,
		time.Second * 5,
		time.Second * 10,
		time.Second * 20,
		time.Second * 40,
		time.Second * 80,
		time.Second * 160,
		time.Second * 360,
		time.Second * 720,
		time.Second * 1440,
	}
	queryRetryTime = []time.Duration{100 * time.Millisecond}
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

type jobConfig struct {
	MaxTimeout time.Duration
	Retires    []time.Duration
}

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
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
			Retires:    defaultRetryTime,
		},
		handler:    handler,
		retryIndex: -1,
		state:      StateInit,
		stopChan:   make(chan bool),
	}
	return &j
}

func NewQueryJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retires:    queryRetryTime,
		},
		handler:    handler,
		retryIndex: -1,
		state:      StateInit,
		stopChan:   make(chan bool),
	}
	return &j
}

func (j *job) Execute(ctx context.Context) error {
	//j.state = StateRunning
	//
	//var err error
	//err = j.handler(ctx)
	//if err != nil {
	//	j.state = StateFailed
	//	return err
	//}
	//
	//j.state = StateCompleted
	//
	//return nil
	ch := make(chan error)
	ctxJob, doneFunc := context.WithCancel(ctx)

	go func() {
		j.state = StateRunning
		var err error

		err = j.handler(ctxJob)

		if err != nil {
			j.state = StateFailed
			ch <- err
			return
		}

		j.state = StateCompleted
		ch <- err
	}()

	//for {
	//	select {
	//	case <-j.stopChan:
	//		break
	//	default:
	//		fmt.Println("Hello world")
	//	}
	//}
	for {
		select {
		case err := <-ch:
			doneFunc()
			return err
		case <-j.stopChan:
			doneFunc()
			return nil
		}
	}

	//return <-ch
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retires[j.retryIndex])

	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retires)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed

	return err
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}
	j.config.Retires = times
}
