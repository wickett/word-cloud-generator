test:
	@go test ./...

run:
	@go run main.go

goconvey:
	$$GOPATH/bin/goconvey

artifact: clean
	@echo "Creating a build as ./artifact/word-cloud-generator"
	@go build -o ./artifact/word-cloud-generator ./main.go

clean:
	@echo "First, cleaning up previous the build"
	@rm -f ./artifact/word-cloud-generator

install:
	@echo "Installs to $$GOPATH/bin"
	@go build ./main.go

git-hooks:
	test -d .git/hooks || mkdir -p .git/hooks
	cp -f hooks/git-pre-commit.hook .git/hooks/pre-commit
	chmod a+x .git/hooks/pre-commit
