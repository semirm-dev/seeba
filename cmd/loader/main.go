package main

import (
	"context"
	"flag"
	"github.com/semirm-dev/seeba/aoe"
	"github.com/semirm-dev/seeba/etl"
)

var (
	importPath = flag.String("i", "data/import/worldofmusic.xml", "path to xml file for import")
	exportPath = flag.String("e", "data/filtered/worldofmusic.xml", "path to export filtered xml file")
	batchSize  = flag.Int("b", 5, "Batch size")
	workers    = flag.Int("w", 5, "Number of data store workers")
)

func main() {
	flag.Parse()

	mainCtx, mainCancel := context.WithCancel(context.Background())
	defer mainCancel()

	ldr := etl.NewLoader(etl.NewFileImporter(*importPath, *batchSize), etl.NewFileExporter(*exportPath), aoe.NewFilter(aoe.NewDefaultValidator()))
	ldr.Load(mainCtx, *workers)
}
