package algorithms

import "fmt"

func BubbleSort(sl []int) []int {
	fmt.Println("----- Bubble Sort -----")
	fmt.Printf("Unordered Slice: \t%v\n", sl)
	// Outer loop controls the number of passes through the slice:
	for pass := 1; pass < len(sl); pass++ {
		// Inner loop performs the comparisons and swaps for each pass
		for i := 0; i < len(sl)-pass; i++ { // the bigger value 'bubbles up' to the last position
			if sl[i] > sl[i+1] { // Swap the elements if the current element is greater than the next
				// Use a temporary variable to hold the value of sl[i]
				temp := sl[i]   // Store sl[i] in temp
				sl[i] = sl[i+1] // Assign sl[i+1] to sl[i]
				sl[i+1] = temp  // Assign temp (original sl[i]) to sl[i+1]
			}
		}
		// After each pass, the largest unsorted element has moved to its correct position
	}
	fmt.Printf("Ordered Slice: \t\t%v\n", sl)
	return sl
}
