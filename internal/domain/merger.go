package domain

type MergerUseCase interface {
	Merge(collection1 []int, collection2 []int, collection3 []int) []int
}
