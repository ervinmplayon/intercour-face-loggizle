package main

import (
	"ervinmplayon/intercour-face-logger/logger"
)

func main() {
	var log logger.Logger = &logger.StandardLogger{}
	log.Info("Live Like Dionysius")
	log.Debug("Imagine Sisyphus Happy")
	log.Error("Other People are Hell")
}
