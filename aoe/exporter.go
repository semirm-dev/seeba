package aoe

import (
	"context"
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

func (exp *exporter) Export(ctx context.Context, musicData []byte) error {
	logrus.Infof("saved music data to file [%s]: %s", exp.dst, musicData)
	return nil
}
