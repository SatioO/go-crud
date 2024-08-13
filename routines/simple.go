package routines

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Producer(ch chan<- int) {
	defer close(ch)

	for i := range 10 {
		ch <- i
	}
}

func Consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		data, ok := <-ch

		if !ok {
			break
		}

		fmt.Printf("consumer: c%d, data:%d\n", id, data)
	}
}

func Buffered(ch chan<- int) {
	defer close(ch)

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
}

func MarkDone(done chan<- bool) {
	time.Sleep(time.Second * 2)
	done <- true
}

func MultiPlex(ch chan<- int) {
	defer close(ch)

	for i := range 10 {
		time.Sleep(time.Second)
		ch <- i
	}
}

func Process(ctx context.Context) {
	for i := range 5 {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Printf("done %d \n", i+1)
		}
	}
}

func Task1(ctx context.Context, wg *sync.WaitGroup) int {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return 0
	case <-time.After(4 * time.Second):
		return 1
	}
}

func Task2(ctx context.Context, wg *sync.WaitGroup) int {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return 0
	case <-time.After(2 * time.Second):
		return 2
	}
}
