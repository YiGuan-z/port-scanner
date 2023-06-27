directory = build
files= mac-arm mac-amd windows-arm windows-amd linux-arm linux-amd

build:all: create-dirs build-mac-arm build-mac-amd build-linux-arm build-linux-amd build-windows-arm build-windows-amd
	@echo "success"
build-mac-amd:
	export $(call env,darwin,amd64) && go build -o $(directory)/mac-amd/port-scanner
build-mac-arm:
	export $(call env,darwin,arm64) && go build -o $(directory)/mac-arm/port-scanner
build-windows-amd:
	export $(call env,windows,amd64) && go build -o $(directory)/windows-amd/port-scanner.exe
build-windows-arm:
	export $(call env,windows,arm64) && go build -o $(directory)/windows-arm/port-scanner.exe
build-linux-amd:
	export $(call env,linux,amd64) && go build -o $(directory)/linux-amd/port-scanner
build-linux-arm:
	export $(call env,linux,arm64) && go build -o $(directory)/linux-arm/port-scanner

create-dirs:
	@if [ -d $(directory) ]; then \
		echo "Directory exists"; \
	else \
		mkdir $(directory); \
		for f in $(files) ; do \
			mkdir $(directory)/$$f; \
		done; \
	fi

clean:
	rm -rf build


env = CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2)