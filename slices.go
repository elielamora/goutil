package goutil

func Map[T any, U any](input []T, mapper func(T) U) []U{
	output := make([]U, len(input))
	for i, x := range input {
		output[i] = mapper(x)
	}
	return output
}