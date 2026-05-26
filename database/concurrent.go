package database

type concurrentOutputWithOrdering struct {
	order  int
	output any
}

func ConcurrentMapFuncWithError[Tin any, Tout any](inputs []Tin, concurrency int, f func(Tin) (Tout, error)) ([]Tout, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// disable concurrency

// no limits
