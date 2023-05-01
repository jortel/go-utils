PKG = ./error/... \
      ./logr/... \
      ./test/...

bin: fmt vet
	go build -o bin/test github.com/jortel/go-utils/test

fmt:
	go fmt ${PKG}

vet:
	go vet ${PKG}

