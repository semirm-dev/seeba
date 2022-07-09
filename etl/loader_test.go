package etl_test

import (
	"context"
	"encoding/json"
	"github.com/semirm-dev/seeba/etl"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockedImporter struct {
	data []string
}

func (imp *mockedImporter) Import(ctx context.Context) *etl.Imported {
	imported := &etl.Imported{
		MusicDataBatch: make(chan []byte),
		OnError:        nil,
	}

	go func() {
		defer close(imported.MusicDataBatch)
		data, _ := json.Marshal(imp.data)
		imported.MusicDataBatch <- data
	}()

	return imported
}

type mockedExporter struct {
	data []string
}

func (exp *mockedExporter) Export(ctx context.Context, filtered []byte) error {
	var data []string
	json.Unmarshal(filtered, &data)
	exp.data = data
	return nil
}

type mockedFilter struct{}

func (flt *mockedFilter) Apply(ctx context.Context, imported []byte) <-chan []byte {
	filtered := make(chan []byte)
	go func() {
		defer close(filtered)
		filtered <- imported
	}()
	return filtered
}

func TestLoader_Load(t *testing.T) {
	testTable := map[string]struct {
		given         []string
		expectedCount int
	}{
		"all data should successfully reach exporter's destination": {
			given: []string{
				"data 1", "data 2", "data 3",
			},
			expectedCount: 3,
		},
	}

	mainCtx, impCancel := context.WithCancel(context.Background())
	defer impCancel()

	for name, suite := range testTable {
		t.Run(name, func(t *testing.T) {
			importer := &mockedImporter{
				data: suite.given,
			}
			exporter := &mockedExporter{}
			filter := &mockedFilter{}

			ldr := etl.NewLoader(importer, exporter, filter)
			ldr.Load(mainCtx, 1)

			assert.Equal(t, suite.expectedCount, len(exporter.data))
		})
	}
}
