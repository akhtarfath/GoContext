package GoContext

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWValue(t *testing.T) {
	contextA := context.Background()    // parent context
	fmt.Println("contextA: ", contextA) // parent context

	contextB := context.WithValue(contextA, "keyB", "valueB") // child context of contextA
	contextD := context.WithValue(contextB, "keyD", "valueD") // child context of contextB
	contextE := context.WithValue(contextB, "keyE", "valueE") // child context of contextB
	fmt.Println("contextB: ", contextB)                       // child context of contextA
	fmt.Println("contextD: ", contextD)                       // child context of contextB
	fmt.Println("contextE: ", contextE)                       // child context of contextB

	contextC := context.WithValue(contextA, "keyC", "valueC") // child context of contextA
	contextF := context.WithValue(contextC, "keyF", "valueF") // child context of contextC
	contextG := context.WithValue(contextF, "keyG", "valueG") // child context of contextC
	fmt.Println("contextC: ", contextC)                       // child context of contextA
	fmt.Println("contextF: ", contextF)                       // child context of contextC
	fmt.Println("contextG: ", contextG)                       // child context of contextC

	fmt.Println("contextF.Value(keyF): ", contextF.Value("keyF")) // child context of contextC
	fmt.Println("contextF.Value(keyC): ", contextF.Value("keyC")) // child context of contextC
	fmt.Println("contextD.Value(keyB): ", contextD.Value("keyB")) // child context of contextC
}
