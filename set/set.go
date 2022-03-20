package set

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(es ...T) {
	for _, e := range es {
		s[e] = struct{}{}
	}
}

func (s Set[T]) Remove(es ...T) {
	for _, e := range es {
		delete(s, e)
	}
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Clear() {
	s = New[T]()
}

func (s Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s Set[T]) Contains(es ...T) bool {
	for _, e := range es {
		_, ok := s[e]

		if !ok {
			return false
		}
	}

	return true
}

func (s Set[T]) Equals(o Set[T]) bool {
	if s.Size() != o.Size() {
		return false
	}

	return s.Contains(s.ToSlice()...)
}

func (s Set[T]) ToSlice() []T {
	slice := []T{}

	for e := range s {
		slice = append(slice, e)
	}

	return slice
}

func (s Set[T]) Clone() Set[T] {
	return New(s.ToSlice()...)
}

func (s Set[T]) ForEach(f func(T)) {
	for e := range s {
		f(e)
	}
}

func New[T comparable](es ...T) Set[T] {
	s := make(Set[T])
	for _, e := range es {
		s.Add(e)
	}

	return s
}
