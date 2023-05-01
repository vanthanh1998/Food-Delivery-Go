package main

import (
	"Food-Delivery/component/asyncjob"
	"context"
	"log"
	"time"
)

func main() {
	//job1 := asyncjob.NewJob(func(ctx context.Context) error {
	//	time.Sleep(time.Second)
	//	log.Println("I am job 1")
	//
	//	//return nil
	//	return errors.New("something went wrong at job 1")
	//})
	//
	//job1.SetRetryDurations([]time.Duration{time.Second * 3})
	//
	//if err := job1.Execute(context.Background()); err != nil {
	//	log.Println(job1.State(), err)
	//
	//	for {
	//		if err := job1.Retry(context.Background()); err != nil {
	//			log.Println(err)
	//		}
	//
	//		if job1.State() == asyncjob.StateRetryFailed || job1.State() == asyncjob.StateCompleted {
	//			log.Println(job1.State())
	//			break
	//		}
	//	}
	//}

	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job 1")

		return nil
		//return errors.New("something went wrong at job 1")
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		log.Println("I am job 2")

		return nil
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 3)
		log.Println("I am job 3")

		return nil
	})

	// isConcurrent = false -> run tuần tự => 1 + 2 +3 = 6s
	// isConcurrent = false -> run đồng thời => 3s
	group := asyncjob.NewGroup(true, job1, job2, job3)

	// func run => use buffer chanel
	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
