package aol

import (
	"context"
	"github.com/semirm-dev/seeba/etl"
)

type importer struct {
}

func NewImporter() *importer {
	return &importer{}
}

func (imp *importer) Import(ctx context.Context) *etl.Imported {
	return nil
}
