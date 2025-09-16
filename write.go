package warppipe

import (
	"encoding/gob"
	"fmt"
	"os"
)

func NewWriter[T any](path string) (*Writer[T], error) {
	if err := createIfNotExists(path); err != nil {
		return nil, err
	}
	w := &Writer[T]{
		open: make(chan struct{}, 1),
	}

	// opening async as 'OpenFile' blocks until the pipe has a reader
	// it's only necessary to block once we need to write
	go func(w *Writer[T]) {
		f, err := os.OpenFile(path, os.O_WRONLY, os.ModeNamedPipe)
		w.f = f
		if err != nil && w.err == nil {
			err = fmt.Errorf("opening [%s]: %w", path, err)
		}
		w.encoder = gob.NewEncoder(f)
		w.open <- struct{}{}
		close(w.open)
	}(w)

	return w, nil
}

type Writer[T any] struct {
	f       *os.File
	encoder *gob.Encoder
	open    chan struct{}
	err     error
}

func (w *Writer[T]) Write(v T) error {
	<-w.open
	if w.err != nil {
		return w.err
	}

	return w.encoder.Encode(v)
}

func (w *Writer[T]) Close() error {
	return w.f.Close()
}
