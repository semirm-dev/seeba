package etl

import (
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// fileImporter will import data from a source file
type fileImporter struct {
	src       string
	batchSize int
}

func NewFileImporter(src string, batchSize int) *fileImporter {
	return &fileImporter{
		src:       src,
		batchSize: batchSize,
	}
}

func (imp *fileImporter) Import(ctx context.Context) *Imported {
	imported := &Imported{
		MusicDataBatch: make(chan []byte),
		OnError:        make(chan error),
	}

	go func() {
		defer close(imported.MusicDataBatch)

		dataFile, err := os.Open(imp.src)
		if err != nil {
			logrus.Fatal(err)
		}
		defer dataFile.Close()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				batch, _ := ioutil.ReadAll(dataFile)

				imported.MusicDataBatch <- batch

				return
			}
		}
	}()

	return imported
}
