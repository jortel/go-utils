GOPATH ?= $(HOME)/go
GOBIN ?= $(GOPATH)/bin
GOIMPORTS = $(GOBIN)/goimports

PKG = ./error/... \
      ./logr/... \
      ./cmd/...

PKGDIR = $(subst /...,,$(PKG))

bin: fmt vet
	go build -o bin/cmd github.com/jortel/go-utils/cmd

fmt: $(GOIMPORTS)
	$(GOIMPORTS) -w $(PKGDIR)

vet:
	go vet ${PKG}

$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports@latest
