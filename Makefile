# .SILENT:

main:
	# Please specify one of the following rules:
	#   gui-mac
	#   clean

gui-mac:
	cd build/gui/macos && make

clean:
	rm -r artifacts