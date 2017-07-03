test:
	@go test $$(go list ./...|grep -v vendor)

run:
	@go run main.go

goconvey:
	$$GOPATH/bin/goconvey

godep:
	@echo "Restoring dependencies..."
	@godep restore

compile: clean godep test
	@echo "Creating cross compiled builds in ./artifacts"
	@env GOOS=darwin GOARCH=amd64 go build -o ./artifacts/osx/word-cloud-generator -v .
	@env GOOS=linux GOARCH=amd64 go build -o ./artifacts/linux/word-cloud-generator -v .
	@env GOOS=windows GOARCH=amd64 go build -o ./artifacts/windows/word-cloud-generator -v .

clean:
	@echo "First, cleaning up the previous build"
	@rm -f ./artifact/word-cloud-generator

install:
	@echo "Installs to $$GOPATH/bin"
	@go build ./main.go

git-hooks:
	test -d .git/hooks || mkdir -p .git/hooks
	cp -f hooks/git-pre-commit.hook .git/hooks/pre-commit
	chmod a+x .git/hooks/pre-commit
