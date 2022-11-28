package main

// Multiple packages can be imported together.
import (
	"fmt"
	"math/rand"
	"time"
)

//driver code
func main() {
	
		// Set the random seed.
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Next pseudo-random number is %v.\n", rand.Uint32()%10)
}