package main

import (
	"offers_iiko/config"
	"offers_iiko/core"
)

func main() {
	core.Start(config.ReleaseMode)
}
