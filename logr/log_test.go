package logr

import (
	"testing"

	liberr "github.com/jortel/go-utils/error"
)

func TestList(t *testing.T) {
	log := WithName("Test")
	type Persion struct {
		Name string
		Age  int
	}
	p := Persion{
		Name: "Elmer",
		Age:  60,
	}

	err := liberr.New("Test")

	log.Info("Test")
	log.Info("Test", "person", p)
	log.Error(err, "Test")
	log.Error(err, "Test", "person", p)
}
