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
