package aol

import (
	"context"
	"github.com/semirm-dev/seeba/etl"
	"github.com/sirupsen/logrus"
)

type exporter struct {
	dst string
}

func NewExporter(dst string) *exporter {
	return &exporter{
		dst: dst,
	}
}

func (exp *exporter) Export(ctx context.Context, musicData []*etl.Music) error {
	for _, d := range musicData {
		logrus.Infof("saved music data to file [%s]: %v", exp.dst, d)
	}
	return nil
}
