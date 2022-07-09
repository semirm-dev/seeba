package etl

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// Importer will import *etl data from its source.
type Importer interface {
	Import(context.Context) *Imported
}

// Exporter will store *etl data in its destination.
type Exporter interface {
	Export(context.Context, []byte) error
}

// Filter will prepare *etl data for writing.
// Provider specific.
type Filter interface {
	Apply(context.Context, []byte) <-chan []byte
}

// Search will get *etl data from its source.
// Provider specific.
type Search interface {
	All() (interface{}, error)
}

// Imported presents each imported *etl data record
type Imported struct {
	MusicDataBatch chan []byte
	OnError        chan error
}

type loader struct {
	importer Importer
	exporter Exporter
	filter   Filter
}

// NewLoader will initialize *loader.
// Loader will load *etl data from Importer, filter it by applying Filter and store it in data store using Exporter
func NewLoader(importer Importer, exporter Exporter, filter Filter) *loader {
	return &loader{
		importer: importer,
		exporter: exporter,
		filter:   filter,
	}
}

// Load will start loading *etl data from Importer to Exporter
func (ldr *loader) Load(ctx context.Context, workers int) {
	t := time.Now()

	logrus.Info("import in progress...")

	imported := ldr.importer.Import(ctx)
	filtered := ldr.filterMusicData(ctx, imported)

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go ldr.exportMusicData(ctx, &wg, filtered)
	}
	wg.Wait()

	logrus.Infof("=== import finished in %v ===", time.Now().Sub(t))
}

// filterMusicData will sanitize and filter *etl data.
func (ldr *loader) filterMusicData(ctx context.Context, imported *Imported) <-chan []byte {
	filtered := make(chan []byte)

	go func() {
		defer close(filtered)

		for {
			select {
			case batch, ok := <-imported.MusicDataBatch:
				if !ok {
					return
				}

				// apply some general business rules, regardless of data provider
				buf := make([]byte, 0)
				buf = append(buf, batch...)

				filterApplied, ok := <-ldr.filter.Apply(ctx, buf)
				if ok {
					filtered <- filterApplied
				}
			case err := <-imported.OnError:
				logrus.Error(err)
			case <-ctx.Done():
				return
			}
		}
	}()

	return filtered
}

// exportMusicData will store *etl data in data store.
// It must be last in the line, all data should already be checked and validated.
func (ldr *loader) exportMusicData(ctx context.Context, wg *sync.WaitGroup, musicData <-chan []byte) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case batch, ok := <-musicData:
			if !ok {
				return
			}

			err := ldr.exporter.Export(ctx, batch)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}
