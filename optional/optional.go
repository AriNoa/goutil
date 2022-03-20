package optional

type Optional[T any] struct {
	value T
	ok    bool
}

func (o Optional[T]) Value() T {
	return o.value
}

func (o Optional[T]) ValueOr(v T) T {
	if o.ok {
		return o.value
	}

	return v
}

func (o Optional[T]) Ok() bool {
	return o.ok
}

func (o Optional[T]) Get() (T, bool) {
	return o.value, o.ok
}

func None[T any]() Optional[T] {
	return Optional[T]{
		ok: false,
	}
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{
		value: value,
		ok:    true,
	}
}
