package model

type PaginationST[T any] struct {
	HasMore bool `json:"hasMore" validate:"required"`
	Items   []T  `json:"items" validate:"required"`
} // @name Pagination
