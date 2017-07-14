BINARY=word-cloud-generator

all: clean godep test build

lint: vet fmt
	@golint $$(go list ./...|grep -v vendor)

vet:
	@go vet $$(go list ./...|grep -v vendor)

fmt:
	@go fmt $$(go list ./...|grep -v vendor)

test:
	@go test $$(go list ./...|grep -v vendor)

run:
	@go run main.go

start-mac: build
	./artifacts/osx/word-cloud-generator

rice:
	@rice embed-go

goconvey-install:
	@go install github.com/smartystreets/goconvey

goconvey:
	$$GOPATH/bin/goconvey -port=9999

godep:
	@echo "Restoring dependencies..."
	@godep restore

build:
	@echo "Creating compiled builds in ./artifacts"
	@env GOOS=darwin GOARCH=amd64 go build -o ./artifacts/osx/${BINARY} -v .
	@env GOOS=linux GOARCH=amd64 go build -o ./artifacts/linux/${BINARY} -v .
	@env GOOS=windows GOARCH=amd64 go build -o ./artifacts/windows/${BINARY} -v .
	@ls -lR ./artifacts

clean:
	@echo "Cleaning up previous builds"
	@go clean
	@rm -rf ./artifacts/*

install:
	@echo "Installs to $$GOPATH/bin"
	@go build ./main.go
	@go install

uninstall:
	@echo "Removing from $$GOPATH/bin"
	@go clean -i

git-hooks:
	test -d .git/hooks || mkdir -p .git/hooks
	cp -f hooks/git-pre-commit.hook .git/hooks/pre-commit
	chmod a+x .git/hooks/pre-commit

.PHONY: all install uninstall clean
.DEFAULT_GOAL := all