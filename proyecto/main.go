package main

import (
	//"capudo/dataBase"
	"capudo/logger"

	"capudo/api"
)

func main() {
	logger.LogInfo("DOJO CAPUDO - welcome!")

	api.Start()
}
