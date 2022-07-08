package aol

import (
	"context"
)

type filter struct {
}

func NewFilter() *filter {
	return &filter{}
}

func (ftr *filter) Apply(ctx context.Context, musicData []byte) <-chan []byte {
	filtered := make(chan []byte)

	go func() {
		defer close(filtered)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				filtered <- musicData
			}
		}
	}()

	return filtered
}
