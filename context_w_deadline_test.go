package GoContext

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContextWithDeadline(t *testing.T) {
	group := sync.WaitGroup{} // create a wait group
	fmt.Println("Num Goroutines: ", runtime.NumGoroutine())

	parent := context.Background()                                             // parent context
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) // child context of parent, with cancel function, deadline is 5 seconds from now
	defer cancel()                                                             // send signal cancel the context, which will close the channel

	destination := CreateCounterDelay(ctx, &group) // create a goroutine that will send data to channel
	fmt.Println("New Num Goroutines: ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter: ", n)
		//if n == 10 {
		//	break
		//}
	}
	group.Wait() // wait for all goroutines to finish
	fmt.Println("Final Num Goroutines: ", runtime.NumGoroutine())
}
