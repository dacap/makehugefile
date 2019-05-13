# Copyright (C) 2019 David Capello

all:
	go build

package:
	env GOOS=darwin GOARCH=386 go build -v -o bin/makehugefile *.go
	cd bin && zip makehugefile-macosx.zip makehugefile && rm makehugefile && cd ..
	env GOOS=windows GOARCH=386 go build -v -o bin/makehugefile.exe *.go
	cd bin && zip makehugefile-windows.zip makehugefile.exe && rm makehugefile.exe && cd ..
	env GOOS=linux GOARCH=386 go build -v -o bin/makehugefile *.go
	cd bin && zip makehugefile-linux.zip makehugefile && rm makehugefile && cd ..
