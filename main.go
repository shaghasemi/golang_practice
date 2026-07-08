package main

import "slices"

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}
	return slices.Min(*numbers)
}

func ReverseSlice(numbers *[]int) {
	slices.Reverse(*numbers)
}

func SwapElements(numbers *[]int, i, j int) {
	if i < 0 || j < 0 || i >= len(*numbers) || j >= len(*numbers) {
		return
	}
	(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
}
