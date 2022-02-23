ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	echo "ROOT_DIR is $(ROOT_DIR)"
    go build -ldflags="-r $(ROOT_DIR)lib" main.go

build-lib: 
	cd lib/crustaceanize && cargo build --release
	cp lib/crustaceanize/target/release/libcrustaceanize.dylib lib/
