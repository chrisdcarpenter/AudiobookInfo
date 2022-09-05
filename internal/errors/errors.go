package errors

import "fmt"

type LoadError struct {
}

func (e LoadError) Error() string {
	return fmt.Sprintf("Failed to load any books")
}
