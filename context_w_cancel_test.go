package GoContext

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func CreateCounter(ctx context.Context, group *sync.WaitGroup) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		group.Add(1)
		counter := 1
		for {
			select {
			case <-ctx.Done(): // if context is canceled, then	break the loop
				group.Done()
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	group := sync.WaitGroup{}
	fmt.Println("Num Goroutines: ", runtime.NumGoroutine())

	parent := context.Background()            // parent context
	ctx, cancel := context.WithCancel(parent) // child context of parent, with cancel function

	destination := CreateCounter(ctx, &group) // create a goroutine that will send data to channel
	fmt.Println("New Num Goroutines: ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter: ", n)
		if n == 10 {
			break
		}
	}
	cancel()     // send signal cancel the context, which will close the channel
	group.Wait() // wait for all goroutines to finish
	fmt.Println("Final Num Goroutines: ", runtime.NumGoroutine())
}
