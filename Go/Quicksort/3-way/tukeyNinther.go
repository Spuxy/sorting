package main

import "fmt"

func TukeyNinther(arr []int) int {
	// var maxMedians = 12
	// var MediansSamples []int
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			fmt.Println("Median vzorek: ", i)
		}
	}
	return 1
}
