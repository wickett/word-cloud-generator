test:
	@go test $$(go list ./...|grep -v vendor)

run:
	@go run main.go

goconvey:
	$$GOPATH/bin/goconvey

godep:
	@echo "Restoring dependencies..."
	@godep restore

all: clean godep test
	@echo "Creating compiled builds in ./artifacts"
	@env GOOS=darwin GOARCH=amd64 go build -o ./artifacts/osx/word-cloud-generator -v .
	@env GOOS=linux GOARCH=amd64 go build -o ./artifacts/linux/word-cloud-generator -v .
	@env GOOS=windows GOARCH=amd64 go build -o ./artifacts/windows/word-cloud-generator -v .
	@ls -lR ./artifacts

clean:
	@echo "Cleaning up the previous build"
	@rm -rf ./artifacts/*

install:
	@echo "Installs to $$GOPATH/bin"
	@go build ./main.go

git-hooks:
	test -d .git/hooks || mkdir -p .git/hooks
	cp -f hooks/git-pre-commit.hook .git/hooks/pre-commit
	chmod a+x .git/hooks/pre-commit

.PHONY: all install clean