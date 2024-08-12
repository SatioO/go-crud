package collections

type Number interface {
}

type List[T any] interface {
	Add(item T)
	Print()
}
