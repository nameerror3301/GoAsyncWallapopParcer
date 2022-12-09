package main

import (
	"GoAsyncWallapopParcer/internal/app"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
}

func main() {
	app.Run()
}
