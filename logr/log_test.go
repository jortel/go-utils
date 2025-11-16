package logr

import (
	"testing"

	liberr "github.com/jortel/go-utils/error"
	"github.com/onsi/gomega"
)

func TestNew(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	log := New("test", 2)
	g.Expect(log).NotTo(gomega.BeNil())
	g.Expect(log.GetSink().(*Sink).level).To(gomega.Equal(2))
}

func TestList(t *testing.T) {
	log := New("Test", 0)
	type Person struct {
		Name string
		Age  int
	}
	p := Person{
		Name: "Elmer",
		Age:  60,
	}

	err := liberr.New("Test")

	log.Info("Test")
	log.Info("Test", "person", p)
	log.Error(err, "Test")
	log.Error(err, "Test", "person", p)
}
