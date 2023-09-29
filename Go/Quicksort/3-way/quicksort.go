package main

import "fmt"

func Quicksort(arr []int) {
	// Insertion sort if array is small subset of items
	if len(arr) <= 20 {
		fmt.Println("Insertion sort")
	}

	// Choose pivot by Tukey's Ninther
	p := TukeyNinther(arr)
	fmt.Println(p)
}
