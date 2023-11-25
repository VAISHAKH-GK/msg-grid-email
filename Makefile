build:
	go build -o ./bin/email ./cmd/

run: build
	./bin/email
