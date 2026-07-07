package main

import (
	"math"
	"strconv"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	sqrt := math.Sqrt(float64(x))
	if sqrt == math.Trunc(sqrt) {
		return true
	}
	return false
}

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var stringifiedNumber string
	var lengthOfNumber int

	stringifiedNumber = strconv.Itoa(x)
	lengthOfNumber = len(stringifiedNumber)

	for i := 0; i < lengthOfNumber/2; i++ {
		if stringifiedNumber[i] != stringifiedNumber[lengthOfNumber-i-1] {
			return false
		}
	}

	return true
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	var result []int
	// for i := 0; i < len(input); i++ {
	// 	if f(input[i]) {
	// 		result = append(result, input[i])
	// 	}
	// }

	for _, value := range input {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

func Map(input []int, m MapperFunc) []int {
	// var result []int
	result := make([]int, 0, len(input))

	for _, value := range input {
		result = append(result, m(value))
	}

	return result
}
