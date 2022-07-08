package aol

import (
	"context"
	"github.com/semirm-dev/seeba/etl"
)

type filter struct {
}

func NewFilter() *filter {
	return &filter{}
}

func (ftr *filter) Apply(ctx context.Context, musicData []*etl.Music) <-chan []*etl.Music {
	filtered := make(chan []*etl.Music)

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
