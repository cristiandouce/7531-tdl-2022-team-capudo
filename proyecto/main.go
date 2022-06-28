package main

import (
	"capudo/logger"
	"capudo/api"
)

func main() {
	logger.LogInfo("DOJO CAPUDO - welcome!")

	api.Start()
}
