package main

import (
	"github.com/ervinmplayon/intercour-face-loggizle/logger"
)

func main() {
	var log logger.Logger = &logger.StandardLogger{}
	log.Info("Live Like Dionysius")
	log.Debug("Imagine Sisyphus Happy")
	log.Error("Other People are Hell")
	// ! what if logged fatal? os.Exit(1) would be called
	// ! ;p

	log = logger.NewLogrusLogger()
	log.Info("Logrus: Live Like Dionysius")
	log.Debug("Logrus: Imagine Sisyphus Happy")
	log.Error("Logrus: Other People are Hell")
}
