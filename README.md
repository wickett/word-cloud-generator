# word-cloud-generator
It's a golang web app that takes a block of text and turns it into a word cloud.

## Notice
This project is under active development. This project is being created as a sample app for an upcoming training class on Continuous Delivery with Lynda.com. You can see previous courses we have made at https://lynda.com/JamesWickett. Thanks!

## Prerequisites
1. Install go - https://golang.org/doc/install (Start learning go with the tour - http://tour.golang.org/)
2. Set $GOPATH `export GOPATH="${HOME}/go"`
3. Set $PATH `export PATH=$PATH:$(go env GOPATH)/bin`
4. Install godep - `go get github.com/tools/godep`
5. Install goconvey - `go get github.com/smartystreets/goconvey`

## Git
We use git hooks to standardize development on the project. Please run `make git-hooks` to get started.

## Test Coverage
`make test`

## Visual Test Coverage
`make goconvey`

## To Run
`make run`
