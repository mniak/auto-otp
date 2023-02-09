# .SILENT:

OSNAME = $(shell uname -s)

build: build-gui
install: install-gui

build-gui: build-gui-$(OSNAME)
install-gui: install-gui-$(OSNAME)

help:
	# Please specify one of the following rules:
	#   build
	#   install
	#   clean

build-gui-Darwin:
	touch cmd/gui/macos
	cd build/gui/macos && $(MAKE) build

install-gui-Darwin:
	touch cmd/gui/macos
	cd build/gui/macos && $(MAKE) install

clean:
	rm -r artifacts
