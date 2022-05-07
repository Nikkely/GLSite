.PHONY: \
	build \
	vendor \
	fmt \
	fmtcheck\
	test \
	rmoutput \

SRC = $(shell git ls-files '*.go')
PKGS = $(shell go list ./...)

build:
	go build -o fetch ./cmd/fetch

test:
	$(foreach pkg,$(PKGS), go test -cover $(pkg) || exit;)

rmoutput:
	rm -rf ./output/*.json

clean:
	echo $(PKGS) | xargs go clean || exit;
