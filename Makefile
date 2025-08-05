GOPATH ?= $(HOME)/go
GOBIN ?= $(GOPATH)/bin
GOIMPORTS = $(GOBIN)/goimports

PKG = ./error/... \
      ./logr/... \
      ./filebacked/... \
      ./cmd/...

PKGDIR = $(subst /...,,$(PKG))

bin: fmt vet
	go build -o bin/cmd github.com/jortel/go-utils/cmd

fmt: $(GOIMPORTS)
	$(GOIMPORTS) -w $(PKGDIR)

vet:
	go vet ${PKG}

test:
	go test -count=1 -v ./error/... ./logr/... ./filebacked/...


$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports@latest
