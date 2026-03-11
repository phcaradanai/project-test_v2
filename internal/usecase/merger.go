package usecase

import "test_v2/internal/domain"

type mergerUseCase struct{}

// NewMergerUseCase creates a new instance of MergerUseCase
func NewMergerUseCase() domain.MergerUseCase {
	return &mergerUseCase{}
}

// Merge combines 3 arrays into one ascending-sorted array.
// collection1 is max to min.
// collection2, collection3 are min to max.
func (u *mergerUseCase) Merge(collection1 []int, collection2 []int, collection3 []int) []int {
	n1 := len(collection1)
	n2 := len(collection2)
	n3 := len(collection3)

	res := make([]int, 0, n1+n2+n3)

	i := n1 - 1 // Start from the end for collection1 (max to min)
	j := 0      // Start from beginning for collection2 (min to max)
	k := 0      // Start from beginning for collection3 (min to max)

	for i >= 0 || j < n2 || k < n3 {
		var val1, val2, val3 int
		has1, has2, has3 := false, false, false

		if i >= 0 {
			val1 = collection1[i]
			has1 = true
		}
		if j < n2 {
			val2 = collection2[j]
			has2 = true
		}
		if k < n3 {
			val3 = collection3[k]
			has3 = true
		}

		// Find the minimum among the available values
		minVal := 0
		minIdx := -1

		if has1 {
			minVal = val1
			minIdx = 1
		}

		if has2 && (minIdx == -1 || val2 < minVal) {
			minVal = val2
			minIdx = 2
		}

		if has3 && (minIdx == -1 || val3 < minVal) {
			minVal = val3
			minIdx = 3
		}

		res = append(res, minVal)

		// Move the pointer of the chosen collection
		if minIdx == 1 {
			i--
		} else if minIdx == 2 {
			j++
		} else if minIdx == 3 {
			k++
		}
	}

	return res
}
