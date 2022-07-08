package etl

import (
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
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
	if err := ioutil.WriteFile(exp.dst, musicData, os.ModePerm); err != nil {
		return err
	}

	logrus.Infof("saved music data to file: %s", exp.dst)

	return nil
}
