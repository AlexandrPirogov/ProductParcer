package comparator

type Comparable[T any] interface {
	GreaterPrice(with T) T
	GreaterRating(with T) T
}
