package lib

import "fmt"

func WrapError(errorMessage string, original error) error {
	return fmt.Errorf("%s\n\tCaused by: %w", errorMessage, original)
}
