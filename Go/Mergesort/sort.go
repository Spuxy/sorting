package main

import (
	"flag"
	"fmt"
)

var (
  DEBUG_MODE = flag.Bool("debug", false, "Shows each step through stdout")
)

func main() {
	flag.Parse()

	inputArr := []int{9, 22, 31, 12, 15, 15, 8, 1, 6, 4, 3, 2}
	fmt.Println("Unsorted array - ", inputArr)
	fmt.Println("Sorted array - ", mergesort(inputArr))

}

func mergesort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	delka := len(arr) / 2
	leftside := arr[:delka]
	rightside := arr[delka:]

	if *DEBUG_MODE == true {
		fmt.Println("Pole: ", arr)
		fmt.Println("Delka poliviny pole: ", delka)
		fmt.Println("Leva strana je: ", leftside)
		fmt.Println("Prava strana je: ", rightside)
	}

	first := mergesort(leftside)   // left side
	second := mergesort(rightside) // right side

	return merge(first, second)
}

func merge(a []int, b []int) []int {
	var finalArr []int
	var j int
	var i int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			// vlozime hodnotu z pole 'a' kam ukazuje index 'i'
			finalArr = append(finalArr, a[i])
			// posuneme ukazatel o dalsi index
			i++
		} else {
			// vlozime hodnotu z pole 'b' kam ukazuje index 'i'
			finalArr = append(finalArr, b[j])
			// posuneme ukazatel o dalsi index
			j++
		}
	}

  // zbytek pole vlozime
	for ; i < len(a); i++ {
		finalArr = append(finalArr, a[i])
	}
  // zbytek pole vlozime
	for ; j < len(b); j++ {
		finalArr = append(finalArr, b[j])
	}

	return finalArr
}
