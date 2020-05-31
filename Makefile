build_date = -X github.com/pconcepcion/telegram_dice_bot/cmd.BuildDate=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
commit_hash = -X github.com/pconcepcion/telegram_dice_bot/cmd.CommitHash=`git rev-parse HEAD`
build:
	go generate ./...
	go build -ldflags "$(build_date) $(commit_hash)" -o dice ./cmd/main.go

install:
	go install

test:
	go test -v -race ./...

gofmt:
	go fmt ./...

lint:
	go vet ./...
	find . -name '*.go' | xargs grep -L '// Code generated by "stringer -type='| xargs golint
	#gocyclo -over 10 .
	errcheck ./...

fuzz: fuzz-build
	mkdir -p go-fuzz-workdir/corpus
	@echo "running go-fuzz, Ctr-C to stop it, next run will start from where it was if the workdir is still intact"
	go-fuzz -bin=dice-fuzz.zip -workdir=go-fuzz-workdir

fuzz-build:
	@echo "Building go-fuzz-build..."
	go-fuzz-build github.com/pconcepcion/dice

clean:
	go clean
	rm -f dice-fuzz.zip

clean-go-fuzz:
	@echo "Cleaning go-fuzz-workdir"
	rm -rf go-fuzz-workdir
	rm -f dice-fuzz.zip
	@echo "Recreating go-fuzz-workdir/corpus"
	mkdir -p go-fuzz-workdir/corpus
deps: dev-deps

dev-deps:
	# go get github.com/golang/lint/golint
	go get -u golang.org/x/lint/golint
	go get github.com/kisielk/errcheck
	go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz-build
