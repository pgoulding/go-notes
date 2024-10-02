package basicexamples

import (
	"errors"
	"fmt"
)

func ExampleErrors() {
	fmt.Println("----- Errors -----")
	n := 10
	findFactorial(n)
	x := -10
	findFactorial(x)

}

// Factorial calculates the factorial of a non-negative integer n.
// It returns 0 as the factorial value if n is negative.
func factorial(n int) (int, error) {

	fmt.Println("Compute factorial of ", n)

	if n < 0 {
		return 0, errors.New("\nError: can not deduce factorial of negative number")
	}

	if n == 0 {
		return 1, nil // Factorial of 0 is 1
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return factorial, nil
}

func findFactorial(fac int) {

	result, err := factorial(fac)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
	} else {
		fmt.Printf("Factorial of %d is %d\n", fac, result)
	}
}
