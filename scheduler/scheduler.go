package scheduler

import (
	"context"
	"sync"
	"time"
)

type JobFunc func()

type Job struct {
	Name     string
	Interval time.Duration
	Func     JobFunc
}

type Scheduler struct {
	jobs   map[string]*Job
	wg     sync.WaitGroup
	cancel context.CancelFunc
}

func New() *Scheduler {
	return &Scheduler{
		jobs: make(map[string]*Job),
	}
}

func (s *Scheduler) AddJob(name string, interval time.Duration, fn JobFunc) {
	s.jobs[name] = &Job{
		Name:     name,
		Interval: interval,
		Func:     fn,
	}
}

func (s *Scheduler) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	for _, job := range s.jobs {
		s.wg.Add(1)
		go s.runJob(ctx, job)
	}
}

func (s *Scheduler) runJob(ctx context.Context, job *Job) {
	defer s.wg.Done()

	ticker := time.NewTicker(job.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			job.Func()
		case <-ctx.Done():
			return
		}
	}
}

func (s *Scheduler) Stop() {
	s.cancel()
	s.wg.Wait()
}
