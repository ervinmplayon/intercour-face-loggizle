package main

import (
	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

func main() {
	var log logger.Logger = &logger.StandardLogger{}
	log.Info("Live Like Dionysius")
	log.Debug("Imagine Sisyphus Happy")
	log.Error("Other People are Hell")

	log = logger.NewLogrusLogger()
	log.Info("Logrus: Live Like Dionysius")
	log.Debug("Logrus: Imagine Sisyphus Happy")
	log.Error("Logrus: Other People are Hell")
}
