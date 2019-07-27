all: vet test testrace

build: deps
	go build github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

clean:
	go clean -i github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

deps:
	go get -d -v github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

test: testdeps
	go test -cpu 1,4 -timeout 7m github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

testdeps:
	go get -d -v -t github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

updatedeps:
	go get -d -v -u -f github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

updatetestdeps:
	go get -d -v -t -u -f github.com/ie310mu/ie310go/forks/google.golang.org/grpc/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps
