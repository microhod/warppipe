//go:build unix

package warppipe

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

func createIfNotExists(path string) error {
	_, err := os.Stat(path)
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := syscall.Mkfifo(path, 0666); err != nil {
		return fmt.Errorf("creating named pipe [%s]: %w", path, err)
	}
	return nil
}
