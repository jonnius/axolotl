package helpers

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetupLogging() error {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
	log.Printf("Starting TextSecure")
	return nil
}
