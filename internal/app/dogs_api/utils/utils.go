package utils

func ReplaceIfNotNil[T any](source *T, destination *T) {
	if source != nil {
		*destination = *source
	}
}
