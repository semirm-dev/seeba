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

func (ftr *filter) Apply(ctx context.Context, imported *etl.Imported) <-chan []*etl.Music {
	return nil
}
