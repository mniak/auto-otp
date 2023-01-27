# .SILENT:

main:
	# Please specify one of the following rules:
	#   app-mac
	#   clean

app-mac:
	touch cmd/gui/macos
	cd build/gui/macos && $(MAKE)

install-mac:
	touch cmd/gui/macos
	cd build/gui/macos && $(MAKE) install

clean:
	rm -r artifacts