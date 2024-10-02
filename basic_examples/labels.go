package basicexamples

import (
	"fmt"
)

func ExampleLabels() {
EXAMPLABEL1: // adding a label for
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				continue EXAMPLABEL1 // jump to the label
			}
			fmt.Printf("i is %d and j is %d\n", i, j)
		}
	}
}
