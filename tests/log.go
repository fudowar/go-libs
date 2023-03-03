package tests

import (
	log "github.com/fudowar/go-libs/log"
)

func TestLog() {
	log.Info("TestLog Info: %s", "test error code")
	log.Error("TestLog Error: %s", "test error code")
	log.Warn("TestLog Warn: %s", "test error code")
}
