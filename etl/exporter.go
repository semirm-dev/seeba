package etl

import (
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

type exporter struct {
	dst      string
	filename string
}

func NewExporter(dst string, filename string) *exporter {
	return &exporter{
		dst:      dst,
		filename: filename,
	}
}

func (exp *exporter) Export(ctx context.Context, musicData []byte) error {
	if _, err := os.Stat(exp.dst); os.IsNotExist(err) {
		if err = os.MkdirAll(exp.dst, os.ModePerm); err != nil {
			return err
		}
	}
	if err := ioutil.WriteFile(path.Join(exp.dst, exp.filename), musicData, os.ModePerm); err != nil {
		return err
	}

	logrus.Infof("saved music data to file: %s", exp.dst)

	return nil
}
