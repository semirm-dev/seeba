package aoe

import (
	"context"
	"github.com/semirm-dev/seeba/etl"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type importer struct {
	src       string
	batchSize int
}

func NewImporter(src string, batchSize int) *importer {
	return &importer{
		src:       src,
		batchSize: batchSize,
	}
}

func (imp *importer) Import(ctx context.Context) *etl.Imported {
	imported := &etl.Imported{
		MusicDataBatch: make(chan []byte),
		OnError:        make(chan error),
	}

	go func() {
		defer close(imported.MusicDataBatch)

		xmlFile, err := os.Open(imp.src)
		if err != nil {
			logrus.Fatal(err)
		}
		defer xmlFile.Close()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				byteValue, _ := ioutil.ReadAll(xmlFile)

				imported.MusicDataBatch <- byteValue

				return
			}
		}
	}()

	return imported
}
