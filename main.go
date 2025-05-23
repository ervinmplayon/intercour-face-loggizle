package main

import (
	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

func main() {
	var log logger.Logger = &logger.StandardLogger{}
	log.Info("Live Like Dionysius")
	log.Debug("Imagine Sisyphus Happy")
	log.Error("Other People are Hell")
}
