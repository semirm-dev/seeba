package main

import (
	"flag"
	"github.com/semirm-dev/seeba/gateway"
	"github.com/semirm-dev/seeba/internal/web"
	"github.com/semirm-dev/seeba/musicmoz"
)

var (
	httpAddr   = flag.String("http", ":8000", "Web server http address")
	exportPath = flag.String("e", "data/filtered/worldofmusic.xml", "path to exported filtered xml file")
)

func main() {
	flag.Parse()

	router := web.NewRouter()

	router.GET("music", gateway.GetMusic(musicmoz.NewSearchApi(*exportPath)))

	web.ServeHttp(*httpAddr, "gateway", router)
}
