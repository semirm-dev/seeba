package exporter

import (
	"context"
	"github.com/semirm-dev/seeba/etl"
	"github.com/sirupsen/logrus"
)

type filesystem struct {
}

func NewFileSystem() *filesystem {
	return &filesystem{}
}

func (exp *filesystem) Export(ctx context.Context, musicData []*etl.Music) error {
	for _, d := range musicData {
		logrus.Infof("saved music data to file: %v", d)
	}
	return nil
}
