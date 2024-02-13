build:
	go build -o ./bin/flash-sale

run: build
	./bin/flash-sale
