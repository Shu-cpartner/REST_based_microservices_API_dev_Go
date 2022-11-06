package main

import (
	"microservicesAPIDevInGolang/app"
	"microservicesAPIDevInGolang/logger"
)

func main() {
	// log.Println("Starting our application...")
	logger.Info("Starting our application...")
	app.Start()
}
