package main

import (
	"flag"
	"offers_iiko/config"
	"offers_iiko/core"
	"offers_iiko/lib/log"
)

func main() {
	var dump string
	//dump  путь куда записывать логи
	flag.StringVar(&dump, "dump", "", "The name of a dump")

	flag.Parse()
	log.PATH = dump

	core.Start(config.ReleaseMode)
}
