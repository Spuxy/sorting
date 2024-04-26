// Bubblesort
//
// Complexity os this sort is `n^2`
// There is `Inner` loop and `Outer` loop
// We need to bubble up higher numbers into the most right side
package main

import "fmt"

func main() {
	fmt.Println("Bubble sorting")
	arr := []int{32, 11, 5, 1, 55, 23, 14, 98, 31, 30, 9, 6}
	fmt.Printf("Array: %d \n", arr)
	Bubblesort(arr)
	fmt.Printf("Result: %d \n", arr)
}

func Bubblesort(arr []int) {
	// Outer loop
	for i := len(arr) - 1; i > 0; i-- {
		// Inner loop
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				// Swap idx
				Swap(arr, j)
			}
		}
	}
}

func Swap(arr []int, idx int) {
	tmpInt := arr[idx]
	arr[idx] = arr[idx+1]
	arr[idx+1] = tmpInt
}
