build:
	go build -o bin/how_rich_am_i

install:
	go install

crosscompile:
	rm -rf bin
	GOOS=linux GOARCH=amd64 go build -o bin/how_rich_am_i_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/how_rich_am_i_darwin_amd64
	GOOS=windows GOARCH=amd64 go build -o bin/how_rich_am_i_windows_amd64
