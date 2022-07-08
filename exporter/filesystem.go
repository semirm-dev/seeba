package exporter

import (
	"context"
	"github.com/semirm-dev/seeba/etl"
)

type filesystem struct {
}

func NewFileSystem() *filesystem {
	return &filesystem{}
}

func (exp *filesystem) Export(ctx context.Context, musicData []*etl.Music) error {
	return nil
}
