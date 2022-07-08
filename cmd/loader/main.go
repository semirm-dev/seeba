package main

import (
	"context"
	"flag"
	"github.com/semirm-dev/seeba/aol"
	"github.com/semirm-dev/seeba/etl"
	"github.com/semirm-dev/seeba/exporter"
)

var (
	xmlPath    = flag.String("p", "data/import/worldofmusic.xml", "path to xml file")
	exportPath = flag.String("e", "data/filtered/worldofmusic.xml", "path to export filtered xml file")
	batchSize  = flag.Int("b", 5, "Batch size")
	workers    = flag.Int("w", 5, "Number of data store workers")
)

func main() {
	flag.Parse()

	impCtx, impCancel := context.WithCancel(context.Background())
	defer impCancel()

	ldr := etl.NewLoader(aol.NewImporter(*xmlPath, *batchSize), aol.NewFilter(), exporter.NewFileSystem(*exportPath))
	ldr.Load(impCtx, *workers)
}
