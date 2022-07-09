package etl

import (
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

// fileExporter will export data to a file destination
type fileExporter struct {
	dest string
}

func NewFileExporter(dst string) *fileExporter {
	return &fileExporter{
		dest: dst,
	}
}

func (exp *fileExporter) Export(ctx context.Context, musicData []byte) error {
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
