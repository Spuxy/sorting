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

	"github.com/Spuxy/sorting/Go/Quicksort/3-way/quicksort"
)

func main() {

	input := []int{99, 4, 6, 2, 5, 8, 13, 66, 1, 3, 100, 32, 12}
	quicksort.Sort(input)
	fmt.Println(input)
}
