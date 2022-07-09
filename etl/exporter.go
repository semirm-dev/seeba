package etl

import (
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

type exporter struct {
	dest string
}

func NewExporter(dst string) *exporter {
	return &exporter{
		dest: dst,
	}
}

func (exp *exporter) Export(ctx context.Context, musicData []byte) error {
	dst := filepath.Dir(exp.dest)

	if _, err := os.Stat(dst); os.IsNotExist(err) {
		if err = os.MkdirAll(dst, os.ModePerm); err != nil {
			return err
		}
	}
	if err := ioutil.WriteFile(exp.dest, musicData, os.ModePerm); err != nil {
		return err
	}

	logrus.Infof("saved music data to file: %s", exp.dest)

	return nil
}
