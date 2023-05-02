PKG = ./error/... \
      ./logr/... \
      ./cmd/...

bin: fmt vet
	go build -o bin/cmd github.com/jortel/go-utils/cmd

fmt:
	go fmt ${PKG}

vet:
	go vet ${PKG}

