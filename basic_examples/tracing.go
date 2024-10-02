package basicexamples

import "fmt"

// entering func
func trace(s string) {
	fmt.Println("Entering: ", s)
}

// leaving func
func untrace(s string) {
	fmt.Println("Leaving: ", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func traceExample() {
	b()
}
