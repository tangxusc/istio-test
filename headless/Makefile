build:
	CGO_ENABLED=0 go build -a -o server cmd/server/server.go
	CGO_ENABLED=0 go build -a -o client cmd/client/client.go
	docker build -f Dockerfile -t headless:test .
	kind load docker-image headless:test --name c1
	rm -rf server client