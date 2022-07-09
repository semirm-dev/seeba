package main

import (
	"context"
	"flag"
	"github.com/semirm-dev/seeba/aoe"
	"github.com/semirm-dev/seeba/etl"
)

var (
	importPath = flag.String("i", "data/import/worldofmusic.xml", "path to xml file for import")
	exportPath = flag.String("e", "data/filtered/", "path to export filtered xml file")
	exportFile = flag.String("f", "worldofmusic.xml", "exported file name")
	batchSize  = flag.Int("b", 5, "Batch size")
	workers    = flag.Int("w", 5, "Number of data store workers")
)

func main() {
	flag.Parse()

	impCtx, impCancel := context.WithCancel(context.Background())
	defer impCancel()

	ldr := etl.NewLoader(etl.NewImporter(*importPath, *batchSize), etl.NewExporter(*exportPath, *exportFile), aoe.NewFilter(aoe.NewDefaultValidator()))
	ldr.Load(impCtx, *workers)
}
