package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit ControlMsg = iota
	ExitOk
)

type Job struct {
	data   int
	result int
}

func doubler(jobs, results chan Job, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("exiting goroutine")
				control <- ExitOk
				return

			default:
				panic("unhandled control message")
			}

		case job := <-jobs:
			results <- Job{data: job.data, result: job.data * 2}

		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	jobs := make(chan Job, 50)
	results := make(chan Job, 50)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{data: i, result: 0}
	}

	for {
		select {
		case res := <-results:
			fmt.Println(res)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit
			<-control
			fmt.Println("program exit")
			return
		}
	}
}
