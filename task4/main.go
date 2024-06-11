package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ch := make(chan int)
	wg := new(sync.WaitGroup)



	fmt.Println("Enter count of workers; \"default\" if you want default count.")
	var userInput string
	if _, err := fmt.Scan(&userInput); err != nil {
		log.Fatalf("Some error in input: %s", err.Error())
	}
	
	workersCount, err := validateInput(userInput)
	if err != nil {
		log.Fatalf("Some error in input-validation: %s", err.Error())
	}

	wp := NewWorkerPool(workersCount)
	wp.SetIds(workersCount)
	fmt.Printf("%d workers are ready for job!\n", workersCount)
	wg.Add(workersCount + 1)

	for i, _ := range wp.Pool {
		go func(w *Worker) {
			defer wg.Done()
			w.Work(ctx, ch)
		}(&wp.Pool[i])
	}

	go func(ctx context.Context, c chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case c <- r.Int():
				fmt.Println("New value in channel was received!")
			case <- ctx.Done():
				defer close(c)
				fmt.Println("Context was cancelled...")
				return

			}

		}
	}(ctx, ch, wg)

	<-ctx.Done()
	wg.Wait()
	for _, worker := range wp.Pool {
		fmt.Printf("Worker %d active status: %t\n", worker.id, worker.active)
	}
	fmt.Println("Go gracefully goes to sleep... z-z-z-z...")

}

func NewWorkerPool(num int) *WorkerPool {
	return &WorkerPool{
		Num: num,
		Pool: make([]Worker, num),
	}
}

type WorkerPool struct {
	Num int
	Pool []Worker
}


func (wp *WorkerPool) SetIds(num int) {
	for i := range wp.Pool {
		wp.Pool[i].id = i + 1
		wp.Pool[i].active = true
	}
}



type Worker struct {
	id int
	active bool
}


func (w *Worker) Work(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("Stopping worker with id %d...\n", w.id)
			w.active = false
			return
		case x, ok := <- ch:
			if !ok {
				w.active = false
				return
			}
			fmt.Println(x)
		}
	}

}


func validateInput(input string) (int, error) {
	if input == "default" {
		return runtime.NumCPU(), nil
	} 
	return strconv.Atoi(input)
}