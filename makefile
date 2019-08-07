all: nl

nl: main.go
	go build -o $@ $^
nl.mac : main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $@ $^
nl.windows: main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $@ $^
nl.linux: main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ $^

clean:
	rm -f nl*
