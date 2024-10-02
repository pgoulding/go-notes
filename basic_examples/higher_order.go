package basicexamples

import (
	"fmt"
)

// aliasing a type
type flt func(int) bool

func isOdd(n int) bool {
	return n%2 != 0
}

func isEven(n int) bool {
	return n%2 == 0
}

// takes in a slice of integers and a function of type flt named above
func filter(sl []int, f flt) []int {
	var res []int
	// iterate through the slices passed as a parameter
	for _, val := range sl {
		// call the function passed and use the resulting boolean for a conditional
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

func ExampleHigherOrder() {
	fmt.Println("----- Structs -----")
	// Pass an int and  the function "Add" as a parameterer to the "callback" function
	callback(1, Add)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)

}

// takes the parameters A and B, and preints the usm of the two
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is %d\n", a, b, a+b)
}

// Receive the int 'y' and the function 'f'
func callback(y int, f func(int, int)) {
	// calls the function passed to it (Add) as 'f' with another number
	f(y, 2)
}
