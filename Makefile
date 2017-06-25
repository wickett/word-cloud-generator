test:
	@go test ./...

run:
	@go run main.go

goconvey:
	$$GOPATH/bin/goconvey

git-hooks:
	test -d .git/hooks || mkdir -p .git/hooks
	cp -f hooks/git-pre-commit.hook .git/hooks/pre-commit
	chmod a+x .git/hooks/pre-commit
