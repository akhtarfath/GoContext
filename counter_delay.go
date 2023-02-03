package GoContext

import (
	"context"
	"sync"
	"time"
)

func CreateCounterDelay(ctx context.Context, group *sync.WaitGroup) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		group.Add(1) // add 1 to wait group to go routine
		counter := 1
		for {
			select {
			case <-ctx.Done(): // if context is canceled, then	break the loop
				group.Done() // done 1 to wait group to go routine, decrease 1
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulate delay
			}
		}
	}()
	return destination
}
