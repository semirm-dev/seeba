package aol

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
		buf := make([]byte, 0)

		defer func() {
			if len(buf) > 0 {
				imported.MusicDataBatch <- buf
			}

			close(imported.MusicDataBatch)
		}()

		xmlFile, err := os.Open(imp.src)
		if err != nil {
			logrus.Fatal(err)
		}
		defer xmlFile.Close()

		byteValue, _ := ioutil.ReadAll(xmlFile)

		imported.MusicDataBatch <- byteValue
	}()

	return imported
}
