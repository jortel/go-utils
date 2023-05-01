package main

import (
	liberr "github.com/jortel/go-utils/error"
	"github.com/jortel/go-utils/logr"
)

func main() {
	var err error
	err = liberr.Wrap(err)
	log := logr.WithName("Test")
	log.Info("Hello World")
}
