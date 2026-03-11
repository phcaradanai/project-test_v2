package usecase

import (
	"reflect"
	"testing"
)

func TestMergerUseCase_Merge(t *testing.T) {
	uc := NewMergerUseCase()

	tests := []struct {
		name string
		col1 []int
		col2 []int
		col3 []int
		want []int
	}{
		{
			name: "All empty",
			col1: []int{},
			col2: []int{},
			col3: []int{},
			want: []int{},
		},
		{
			name: "Normal case",
			col1: []int{9, 7, 5, 3, 1}, // max to min
			col2: []int{0, 2, 4, 6, 8}, // min to max
			col3: []int{0, 5, 10},      // min to max
			want: []int{0, 0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "With negatives",
			col1: []int{5, 0, -5}, // max to min
			col2: []int{-10, 2},   // min to max
			col3: []int{-8, -2},   // min to max
			want: []int{-10, -8, -5, -2, 0, 2, 5},
		},
		{
			name: "Only collection1 has items",
			col1: []int{5, 4, 3, 2, 1},
			col2: []int{},
			col3: []int{},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Only collection2 has items",
			col1: []int{},
			col2: []int{1, 2, 3, 4, 5},
			col3: []int{},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uc.Merge(tt.col1, tt.col2, tt.col3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
