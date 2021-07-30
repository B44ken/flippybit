.PHONY: run build

build: 
	[ -d build ] || mkdir build
	touch build/dummy
	rm build/*
	export GOOS='linux'
	go build -o build/flippybit ./src

run: 
	go run ./src
