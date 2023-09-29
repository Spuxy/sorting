package pivot

import "fmt"

const maxMedians = 12

func TukeyNinther(arr []int) int {
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
