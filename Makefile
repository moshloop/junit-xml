.PHONY: build
build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	mkdir -p build
	gox -os="darwin linux windows" -arch="amd64"
	mkdir -p build
	mv junit-xml_darwin_amd64 build/junit-xml_osx
	mv junit-xml_linux_amd64 build/junit-xml
	mv junit-xml_windows_amd64.exe build/junit-xml.exe