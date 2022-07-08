package exporter

import (
	"context"
	"github.com/semirm-dev/seeba/etl"
	"github.com/sirupsen/logrus"
)

type filesystem struct {
	dst string
}

func NewFileSystem(dst string) *filesystem {
	return &filesystem{
		dst: dst,
	}
}

func (exp *filesystem) Export(ctx context.Context, musicData []*etl.Music) error {
	for _, d := range musicData {
		logrus.Infof("saved music data to file [%s]: %v", exp.dst, d)
	}
	return nil
}
