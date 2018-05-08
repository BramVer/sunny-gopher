# sunny-gopher Makefile
#
# This program is free software; you can redistribute
# it and/or modify it under the terms of the GNU
# General Public License …

SHELL = /bin/sh

srcdir = .

NAME 		= sunny-gopher
DESCRIPTION = Returns a random quote from IASIP.
VERSION 	= ${SUNNY_GOPHER_VERSION}
ARCH 		= x86_64

all: build

build:
	docker run --rm -v "$$PWD":/go/src/sunny-gopher -w /go/src/sunny-gopher golang:1.10 go build

TMPDIR := $(shell mktemp -d)
TARGET := $(TMPDIR)/opt/sunny-gopher
CONFIG := $(TARGET)/config
SYSTEM := $(TMPDIR)/usr/lib/systemd/system/

package:
	mkdir -p $(CONFIG)
	mkdir -p $(SYSTEM)

	strip sunny-gopher
	cp ./sunny-gopher $(TARGET)/
	cp ./sunny-gopher.service $(SYSTEM)/sunny-gopher.service
	
	fpm -s dir -t rpm \
		--name "$(NAME)" \
		--description "$(DESCRIPTION)" \
		--version "$(VERSION)" \
		--architecture "$(ARCH)" \
		--iteration $(BUILD_NO) \
		--force \
		--config-files /usr/lib/systemd/system/sunny-gopher.service \
		--chdir $(TMPDIR) \
		.; \
	
	rm -R $(TMPDIR)

clean:
	rm -f sunny-gopher*.rpm

.PHONY: clean
