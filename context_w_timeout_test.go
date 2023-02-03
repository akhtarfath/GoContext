package GoContext

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContextWithTimeOut(t *testing.T) {
	group := sync.WaitGroup{} // create a wait group
	fmt.Println("Num Goroutines: ", runtime.NumGoroutine())

	parent := context.Background()                             // parent context
	ctx, cancel := context.WithTimeout(parent, 10*time.Second) // child context of parent, with cancel function
	defer cancel()                                             // send signal cancel the context, which will close the channel

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
