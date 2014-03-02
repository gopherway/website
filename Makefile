SOURCE = $(wildcard *.go)
TAG = $(shell git describe --tags)
GOBUILD = go build -ldflags '-w'

# $(tag) here will contain either `-1.0-` or just `-`
ALL = \
	$(foreach arch,32 64,\
	$(foreach tag,-$(TAG)- -,\
	$(foreach suffix,win.exe linux osx,\
		build/gopherway$(tag)$(arch)-$(suffix))))

run:
	go run *.go

all: $(ALL)

clean:
	rm -f $(ALL)

fmt:
	gofmt -w=true *.go

# os is determined as thus: if variable of suffix exists, it's taken, if not, then
# suffix itself is taken
win.exe = windows
osx = darwin
build/gopherway-$(TAG)-64-%: $(SOURCE)
	# cd $(GOROOT)/src/; sudo CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=amd64 ./make.bash --no-clean
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=amd64 $(GOBUILD) -o $@

build/gopherway-$(TAG)-32-%: $(SOURCE)
	# cd $(GOROOT)/src/; sudo CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=386 ./make.bash --no-clean
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=386 $(GOBUILD) -o $@

build/gopherway-%: build/gopherway-$(TAG)-%
	@mkdir -p $(@D)
	cd $(@D) && ln -sf $(<F) $(@F)

upload: $(ALL)
ifndef UPLOAD_PATH
	@echo "Define UPLOAD_PATH to determine where files should be uploaded"
else
	rsync -l -P $(ALL) $(UPLOAD_PATH)
endif