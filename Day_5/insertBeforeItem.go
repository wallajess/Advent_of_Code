package main

import "slices"

func insertBefore(slice []int, index1, index2 int) []int {
	toFirst := slice[index2]
	slice = slices.Delete(slice, index2, index2+1)
	slice = slices.Insert(slice, index1, toFirst)

	return slice
}
