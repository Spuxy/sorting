package quicksort

import (
	"github.com/Spuxy/sorting/Go/Quicksort/3-way/partition"
	"github.com/Spuxy/sorting/Go/Quicksort/3-way/pivot"
)

func Sort(v []int) {
	if len(v) < 5 {
		InsertionSort(v)
		return
	}
	p := pivot.TukeyNinther(v)
	low, high := partition.ThreeWay(v, p)
	Sort(v[:low])
	Sort(v[high:])
}

func InsertionSort(v []int) {
	for j := 1; j < len(v); j++ {
		// Invariant: v[:j] contains the same elements as
		// the original slice v[:j], but in sorted order.
		key := v[j]
		i := j - 1
		for i >= 0 && v[i] > key {
			v[i+1] = v[i]
			i--
		}
		v[i+1] = key
	}
}
