package GoContext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	// context.Background() returns a non-nil, empty Context.
	//It is never canceled, has no values, and has no deadline.
	//It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
	background := context.Background()
	fmt.Println("Background: ", background)

	// context.TODO() returns a non-nil, empty Context.
	// Code should use context.TODO() when it's unclear which Context to use or it is not yet available
	// (because the surrounding function has not yet been extended to accept a Context parameter).
	todo := context.TODO()
	fmt.Println("Todo: ", todo)
}
