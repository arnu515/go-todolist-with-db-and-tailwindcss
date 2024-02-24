package util

func Must[T any](a T, err any) T {
	if err != nil {
		panic(err)
	}
	return a
}
