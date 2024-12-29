package asyncjob

import (
	"context"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"log"
	"sync"
)

type group struct {
	jobs         []Job
	isConcurrent bool
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	g := &group{
		jobs:         jobs,
		isConcurrent: isConcurrent,
		wg:           new(sync.WaitGroup),
	}
	return g
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	for i, _ := range g.jobs {
		if g.isConcurrent {
			go func(aj Job) {
				defer core.AppRecover()
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(g.jobs[i])
			continue
		}

		job := g.jobs[i]
		errChan <- g.runJob(ctx, job)
		g.wg.Done()
	}

	var err error

	g.wg.Wait()

	for i := 1; i < len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			return v
		}
	}

	return err

}

// Retry if needed
func (g *group) runJob(ctx context.Context, j Job) error {
	if err := j.Execute(ctx); err != nil {
		for {
			log.Println(err)
			if j.State() == StateRetryFailed {
				return err
			}

			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}
	return nil
}
