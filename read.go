package warppipe

import (
	"encoding/gob"
	"fmt"
	"os"
)

func NewReader[T any](path string) (*Reader[T], error) {
	if err := createIfNotExists(path); err != nil {
		return nil, err
	}
	r := &Reader[T]{
		open: make(chan struct{}, 1),
	}

	// opening async as 'OpenFile' blocks until the pipe has a writer
	// it's only necessary to block once we need to read
	go func(r *Reader[T]) {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModeNamedPipe)
		if err != nil {
			r.err = fmt.Errorf("opening [%s]: %w", path, err)
		}
		r.f = f
		r.decoder = gob.NewDecoder(f)
		r.open <- struct{}{}
		close(r.open)
	}(r)

	return r, nil
}

type Reader[T any] struct {
	f       *os.File
	decoder *gob.Decoder
	open    chan struct{}
	err     error
}

func (r *Reader[T]) Read(v *T) error {
	<-r.open
	if r.err != nil {
		return r.err
	}

	return r.decoder.Decode(v)
}

func (r *Reader[T]) Close() error {
	return r.f.Close()
}
