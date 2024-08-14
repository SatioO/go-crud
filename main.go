package main

import (
	"fmt"
	"sync"
)

type Job struct {
	id  int
	num int
}

type Result struct {
	job Job
	sum int
}

func allocate(jobs chan<- Job) {
	for i := range 10000 {
		job := Job{id: i, num: i * 765}
		jobs <- job
	}

	close(jobs)
}

func process(i int, jobs <-chan Job, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Printf("processor: %d, recieved: %d\n", i, job.num)
	}

	wg.Done()
}

func createWorkerPool(jobs <-chan Job) {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go process(i, jobs, &wg)
	}

	wg.Wait()
}

// func delay(cancel context.CancelFunc) {
// 	time.Sleep(time.Second * 4)
// 	cancel()
// }

func main() {
	// Environment
	// fmt.Println(os.Getenv("DB_CONNECTION"))
	// Heartbeat
	// ch := make(chan string, 1)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)

	// go routines.HeartBeat(ch, ctx, time.Second)
	// go delay(cancel)

	// for {
	// 	_, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}

	// 	fmt.Println("heartbeat")
	// }

	// WORKER POOL
	// Worker goroutines listens for new tasks on jobs buffered channel,
	// once the job is completed, result is written to the results buffered channel
	// jobs := make(chan Job, 10)
	// results := make(chan Result, 10)

	// go allocate(jobs)
	// // go process(jobs)

	// createWorkerPool(jobs)
	// BUFFERED
	// ch := make(chan int, 2)
	// go routines.Buffered(ch)

	// time.Sleep(2 * time.Second)
	// for v := range ch {
	// 	fmt.Println("read value", v, "from ch")
	// 	time.Sleep(2 * time.Second)
	// }

	//Single Producer Multiple Consumer
	// ch := make(chan int)
	// var wg sync.WaitGroup

	// go routines.Producer(ch)

	// wg.Add(1)
	// go routines.Consumer(1, ch, &wg)
	// wg.Add(1)
	// go routines.Consumer(2, ch, &wg)

	// wg.Wait()

	// EXAMPLE 6
	// ch := make(chan int)
	// go routines.Buffered(ch)

	// for {
	// 	data, ok := <-ch

	// 	if !ok {
	// 		break
	// 	}

	// 	fmt.Println(data)
	// }
	// EXAMPLE 5
	// ch := make(chan int, 2)
	// go routines.Buffered(ch)

	// for i := range ch {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(i)
	// }

	// EXAMPLE 4
	// done := make(chan bool, 1)
	// go routines.MarkDone(done)
	// <-done

	// EXAMPLE 3
	// ch := make(chan int)

	// go routines.MultiPlex(ch)

	// for i := range ch {
	// 	fmt.Println(i)
	// }

	// EXAMPLE 2
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// ch1 := make(chan int, 1)
	// ch2 := make(chan int, 1)

	// defer cancel()

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	ch1 <- routines.Task1(ctx, &wg)
	// }()

	// wg.Add(1)
	// go func() {
	// 	ch2 <- routines.Task2(ctx, &wg)
	// }()

	// wg.Wait()

	// output := <-ch1 + <-ch2
	// fmt.Println(output)
}
