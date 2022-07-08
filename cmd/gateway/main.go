package main

import (
	"flag"
	"github.com/semirm-dev/seeba/gateway"
	"github.com/semirm-dev/seeba/internal/web"
)

var (
	httpAddr = flag.String("http", ":8000", "Web server http address")
)

func main() {
	flag.Parse()

	router := web.NewRouter()

	router.GET("music", gateway.GetMusic(nil))

	web.ServeHttp(*httpAddr, "gateway", router)
}
