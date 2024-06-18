package utils

import (
	"fmt"
)

func IsArraySame[T comparable](numsA, numsB []T) bool {
	if len(numsA) != len(numsB) {
		fmt.Printf("The length is not the same (%d != %d)", len(numsA), len(numsB))
		return false
	}
	for i := 0; i < len(numsA); i++ {
		if numsA[i] != numsB[i] {
			fmt.Printf("In index %d, item %v is not the same with %v", i, numsA[i], numsB[i])
			return false
		}
	}
	return true
}

func Is2DArraySame[T comparable](numsA, numsB [][]T) bool {
	if len(numsA) != len(numsB) {
		fmt.Printf("The length of 1st level array is not the same (%d != %d)", len(numsA), len(numsB))
		return false
	}
	for i := 0; i < len(numsA); i++ {
		if len(numsA[i]) == len(numsB[i]) {
			for j := 0; j < len(numsA[i]); j++ {
				if numsA[i][j] != numsB[i][j] {
					fmt.Printf("In index [%d][%d], item %v is not the same with %v", i, j, numsA[i][j], numsB[i][j])
					return false
				}
			}
		} else {
			fmt.Printf("The length of 2nd level array in index %d is not the same (%d != %d)", i, len(numsA[i]), len(numsB[i]))
			return false
		}
	}
	return true
}
