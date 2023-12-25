package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	newjob := cron.New()

	minute := 29
	hour := 10
	day := 25
	month := 12

	entryId1, _ := newjob.AddFunc(fmt.Sprintf("%d %d %d %d *", minute, hour, day, month), func() {
		fmt.Println("Job 1")
	})

	entryId2, _ := newjob.AddFunc(fmt.Sprintf("%d %d %d %d *", minute+1, hour, day, month), func() {
		fmt.Println("Job 2")
	})

	newjob.Start()

	newjob.Remove(entryId1)

	entryId3, _ := newjob.AddFunc(fmt.Sprintf("%d %d %d %d *", minute, hour, day, month), func() {
		fmt.Println("Job 3")
	})

	fmt.Println(entryId1, entryId2, entryId3)

	// Run the scheduler for a duration to allow the jobs to execute
	// For demonstration purposes, this example runs for 1 minute
	time.Sleep(3 * time.Minute)

	// Stop the cron scheduler (optional)
	newjob.Stop()
}
