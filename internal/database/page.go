package database

type PageResult[T any] struct {
	Items []T
	Total int64
}
