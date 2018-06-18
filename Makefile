.PHONY: build
build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	mkdir -p build
	gox -os="darwin linux windows" -arch="amd64"
	mkdir -p build
	mv junit-xml_darwin_amd64 build/waiter_osx
	mv junit-xml_linux_amd64 build/waiter
	mv junit-xml_windows_amd64.exe build/waiter.exe