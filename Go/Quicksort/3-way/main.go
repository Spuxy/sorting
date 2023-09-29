// Quicksort
// Choosing pivot schema:
// This sort will use 3-way partitioning
// That means We need choose carefully pivot
// Neverlessly We goin to use Turkey's Ninther
//
// Choosing sort schema for small subset:
// We need to use simpler algorithm `insertion sort`
package main

import (
	"fmt"
)

func main() {
	input := []int{4, 2, 5, 8, 13, 66, 1, 3}
	Quicksort2(input)
}

func Quicksort2(arr []int) {
	// Insertion sort if array is small subset of items
	if len(arr) <= 5 {
		fmt.Println("Insertion sort")
	}

	// Choose pivot by Tukey's Ninther
	fmt.Println(arr)
	p := TukeyNinther2(arr)
	low, high := Partition(arr, p)
	Quicksort2(arr[:low])
	Quicksort2(arr[high:])
}

const maxMedians = 12

func TukeyNinther2(arr []int) int {
	var medianCounter int
	var groupsSamples []int
	var medianSamples []int
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			if medianCounter >= maxMedians {
				fmt.Println("Can not be more than 12 samples!")
				break
			}
			medianCounter++
			groupsSamples = append(groupsSamples, arr[i])
		}
	}
	fmt.Println(" TUKEY ")
	fmt.Println(groupsSamples)
	fmt.Println(medianSamples)
	if len(groupsSamples) <= 1 {
		return groupsSamples[0]
	}
	for k, v := range groupsSamples {
		if k%2 == 1 {
			medianSamples = append(medianSamples, v)
		}
	}
	return medianSamples[0]
}

func Partition(v []int, p int) (int, int) {
	// Reorder/partition of `arr`
	// Where elements less than `pivot` will be on left side
	// Where elements greater than `pivot` will be on right side
	// Where elements equals to `pivot` will be in the middle
	low, high := 0, len(v)
	for mid := 0; mid < high; {
		// Invariant:
		//  - v[:low] < p
		//  - v[low:mid] = p
		//  - v[mid:high] are unknown
		//  - v[high:] > p
		//
		//         < p       = p        unknown       > p
		//     -----------------------------------------------
		// v: |          |          |a            |           |
		//     -----------------------------------------------
		//                ^          ^             ^
		//               low        mid           high
		switch a := v[mid]; {
		case a < p:
			v[mid] = v[low]
			v[low] = a
			low++
			mid++
		case a == p:
			mid++
		default: // a > p
			v[mid] = v[high-1]
			v[high-1] = a
			high--
		}
	}
	return low, high
}
