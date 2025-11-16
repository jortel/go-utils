package main

import (
	"errors"

	liberr "github.com/jortel/go-utils/error"
	"github.com/jortel/go-utils/logr"
)

func main() {
	log := logr.New("Test", 0, "and", "fields")

	log.Info("HELLO", "name", "jeff")

	err := liberr.New("This failed.", "name", "elmer")
	log.Error(err, "Test this error message.")

	err = liberr.Wrap(err, "This is bad.")
	log.Error(err, "Test this error message (2).")

	err = liberr.Wrap(err, "Wrapped again.")
	log.Error(err, "Test this error message (3).")

	err = errors.New("plain error")
	log.Error(err, "")
}
