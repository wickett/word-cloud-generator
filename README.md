[![Go Test](https://github.com/wickett/word-cloud-generator/actions/workflows/go.yml/badge.svg)](https://github.com/wickett/word-cloud-generator/actions/workflows/go.yml) [![CircleCI](https://circleci.com/gh/wickett/word-cloud-generator/tree/master.svg?style=svg)](https://circleci.com/gh/wickett/word-cloud-generator/tree/master)

# word-cloud-generator
It's a golang web app that takes a block of text and turns it into a word cloud. 
Star this repo if you like the class!

## Notice
This project is a a sample app for a class on Continuous Delivery at LinkedIn Learning. You can see the class as well as other classes I have available at https://www.linkedin.com/learning/instructors/james-wickett. Thanks!

## Prerequisites
1. Install go - https://golang.org/doc/install (Start learning go with the tour - http://tour.golang.org/)
2. Set $GOPATH `export GOPATH="${HOME}/go"`
3. Set $PATH `export PATH=$PATH:$(go env GOPATH)/bin`
5. Install goconvey - `go get github.com/smartystreets/goconvey`

## Git
We use git hooks to standardize development on the project. Please run `make git-hooks` to get started.

## Using Make

### Building Artifacts
This will pull down dependencies, run unit tests, and compile a linux, mac and windows binary into ./artifacts.

`make`

### Fiddling Around

You can build and install a copy in your local $GOPATH/bin directory with:

```
make install
```

### Clean up previous build files
```
make clean
```

### Linting
```
make lint
```

### Test Coverage
```
make test
```

### Visual Test Coverage
```
make goconvey-install
make goconvey
```

### To Run
`make run` or just run the executable.  It will run as a daemon and bind to port 8888, and you can see it by going to http://localhost:8888 in your browser.

## Use API via Curl
```
$ curl -H "Content-Type: application/json" -d '{"text":"ths is a really really really important thing this is"}' http://localhost:8888/api
```

## Setup example
Use this setup

aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

