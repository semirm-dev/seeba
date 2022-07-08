package aol

import (
	"context"
	"fmt"
	"github.com/semirm-dev/seeba/etl"
)

type importer struct {
	batchSize int
}

func NewImporter(s string, batchSize int) *importer {
	return &importer{
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

		for i := 0; i < imp.batchSize*2; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				buf = append(buf, &etl.Music{
					Name:       fmt.Sprintf("music %d", i),
					TrackCount: i,
				})

				if len(buf) >= imp.batchSize {
					imported.MusicDataBatch <- buf
					buf = nil
				}
			}
		}
	}()

	return imported
}
