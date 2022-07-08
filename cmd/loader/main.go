package main

import (
	"context"
	"flag"
	"github.com/semirm-dev/seeba/aol"
	"github.com/semirm-dev/seeba/etl"
	"github.com/semirm-dev/seeba/exporter"
)

var (
	xmlPath = flag.String("p", "cmd/loader/worldofmusic.xml", "path to xml file")
	batch   = flag.Int("b", 400, "Batch size")
	workers = flag.Int("w", 5, "Number of data store workers")
)

func main() {
	flag.Parse()

	impCtx, impCancel := context.WithCancel(context.Background())
	defer impCancel()

	ldr := etl.NewLoader(aol.NewImporter(), aol.NewFilter(), exporter.NewFileSystem())
	ldr.Load(impCtx, *workers)
}
