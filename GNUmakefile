.PHONY: \
	build \
	vendor \
	fmt \
	fmtcheck\
	vet \
	test \
	rmoutput \

SRCS = $(shell git ls-files '*.go')
PKGS = $(shell go list ./...)

build:
	go build -o GLSite .

test:
	$(foreach pkg,$(PKGS), go test -cover $(pkg) || exit;)

rmoutput:
	rm -rf ./output/*.json

vet:
	echo $(PKGS) | xargs env go vet || exit;

fmt:
	gofmt -s -w $(SRCS)

fmtcheck:
	$(foreach file,$(SRCS),gofmt -s -d $(file);)

clean:
	echo $(PKGS) | xargs go clean || exit;
