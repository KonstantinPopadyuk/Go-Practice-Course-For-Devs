package task09

type Set[T comparable] struct {
	items map[T]struct{}
}

func (s *Set[T]) Add(value T) {
	// TODO: implement.
}

func (s *Set[T]) Remove(value T) {
	// TODO: implement.
}

func (s Set[T]) Contains(value T) bool {
	// TODO: implement.
	return false
}

func (s Set[T]) Len() int {
	// TODO: implement.
	return 0
}

func (s Set[T]) Values() []T {
	// TODO: implement.
	return nil
}

func (s Set[T]) Clone() Set[T] {
	// TODO: implement.
	return Set[T]{}
}
