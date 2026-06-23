package main

import (
	"fmt"
	"time"

	"github.com/LodiAleardo/go-cron-scheduler/scheduler"
)

func main() {
	s := scheduler.New()

	s.AddJob("job-1", 2*time.Second, func() {
		fmt.Printf("[%s] job-1 running\n", time.Now().Format("15:04:05"))
	})

	s.AddJob("job-2", 3*time.Second, func() {
		fmt.Printf("[%s] job-2 running\n", time.Now().Format("15:04:05"))
	})

	s.AddJob("job-3", 5*time.Second, func() {
		fmt.Printf("[%s] job-3 running\n", time.Now().Format("15:04:05"))
	})

	s.Start()
	time.Sleep(12 * time.Second)
	s.Stop()
	fmt.Println("scheduler stopped")
}
