package lib

import (
	"fmt"
	"io"
	"log/slog"
)

func WrapError(errorMessage string, original error) error {
	return fmt.Errorf("%s\n\tCaused by: %w", errorMessage, original)
}

func CloseOrWarn(closer io.Closer) {
	if err := closer.Close(); err != nil {
		slog.Warn("error while closing a Closer: " + err.Error())
	}
}
