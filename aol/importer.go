package aol

import (
	"context"
	"encoding/xml"
	"fmt"
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
		MusicDataBatch: make(chan []*etl.Music),
		OnError:        make(chan error),
	}

	go func() {
		buf := make([]*etl.Music, 0)

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

		var data *struct {
			XMLName xml.Name `xml:"records"`
			Records []*struct {
				Title        string `xml:"title"`
				ReleaseDate  string `xml:"releasedate"`
				TrackListing []*struct {
					Track string `xml:"track"`
				} `xml:"tracklisting"`
			} `xml:"record"`
		}
		if err = xml.Unmarshal(byteValue, &data); err != nil {
			logrus.Fatal(err)
		}

		buf = append(buf, &etl.Music{
			Name:       fmt.Sprintf("music 1"),
			TrackCount: 1,
		})

		if len(buf) >= imp.batchSize {
			imported.MusicDataBatch <- buf
			buf = nil
		}
	}()

	return imported
}
