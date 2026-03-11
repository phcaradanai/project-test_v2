package main

import (
	"fmt"
	"test_v2/internal/usecase"
)

func main() {
	merger := usecase.NewMergerUseCase()

	col1 := []int{9, 7, 5, 3, 1}
	col2 := []int{0, 2, 4, 6, 8}
	col3 := []int{0, 5, 10}

	result := merger.Merge(col1, col2, col3)

	fmt.Println("Collection 1:", col1)
	fmt.Println("Collection 2:", col2)
	fmt.Println("Collection 3:", col3)
	fmt.Println("Merged Result:", result)
}
