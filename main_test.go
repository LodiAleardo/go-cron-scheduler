package main

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/LodiAleardo/go-cron-scheduler/scheduler"
)

func TestScheduler_RunsJobs(t *testing.T) {
	s := scheduler.New()
	var counter int32

	s.AddJob("test-job", 50*time.Millisecond, func() {
		atomic.AddInt32(&counter, 1)
	})

	s.Start()
	time.Sleep(250 * time.Millisecond)
	s.Stop()

	got := atomic.LoadInt32(&counter)
	if got < 4 {
		t.Errorf("expected at least 4 runs, got %d", got)
	}
}
