# To build several target concurrently executes make with -j n argument, where n is the number of simultaneous jobs.
GO_EXECUTABLE ?= go
REMOTE_ACCESS_BASENAME ?= remote-access
# Extract the first path of the GOPATH env var
REMOTE_ACCESS_GOPATH := $(shell echo $$GOPATH | cut -d ':' -f 2)
REMOTE_ACCESS_VERSION ?= $(shell git describe --tags)
REMOTE_ACCESS_SG ?= undefined
REMOTE_ACCESS_REGION ?= undefined

.PHONY: release build $(PLATFORMS) clean

# Thank you Vic: https://vic.demuzere.be/articles/golang-makefile-crosscompile/
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

release: build

build: $(PLATFORMS)
$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o $(REMOTE_ACCESS_BASENAME)-$(os)-$(arch)-$(REMOTE_ACCESS_VERSION) -ldflags "-X main.version=$(REMOTE_ACCESS_VERSION) -X main.securitygroup=$(REMOTE_ACCESS_SG) -X main.region=$(REMOTE_ACCESS_REGION)"

clean:
	rm -f $(REMOTE_ACCESS_BASENAME)-*
	rm -f ${REMOTE_ACCESS_GOPATH}/bin/$(REMOTE_ACCESS_BASENAME)-*


