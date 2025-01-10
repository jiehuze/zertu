DATE = $(shell date  +%Y年%m月%d日-%H:%M:%S)
USER = $(shell whoami)

VERSION_INFO = $(shell git tag | awk 'END {print}')
APP = $(shell pwd | awk -F '/' '{print $$NF}')

BUILD_INFO = [$(USER)]$(DATE)
LDFLAGS = -ldflags "-X main.BuildInfo=$(BUILD_INFO)	\
				          -X main.VersionInfo=$(VERSION_INFO)"

all: clean init $(APP) tarball

init:
	mkdir -p target

$(APP):
	@echo "building $(APP) ..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o target/$(APP)  main.go

mac: clean init
	@echo "building ios $(APP) ..."
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o target/$(APP)  main.go

riscv: clean init
	@echo "building ios $(APP) ..."
	GOOS=linux GOARCH=riscv64 go build $(LDFLAGS) -o target/$(APP)  main.go

tarball:
	cp -rf conf target/
	#cp -rf runtime/startup.sh target/
#	cd target/ && tar -zcf $(APP)-$(VERSION_INFO).tar.gz conf $(APP)

clean:
	rm -rf $(APP)
	rm -rf target
