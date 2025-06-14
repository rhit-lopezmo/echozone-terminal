package utils

func Map[T, V any](elements []T, callback func(T) V) []V {
    result := make([]V, len(elements))

		for i, element := range elements {
        result[i] = callback(element)
    }

    return result
}

func MapNoReturn[T any](elements []T, callback func(T)) {
		for _, element := range elements {
       callback(element)
    }
}

