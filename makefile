build:
	go build -o ./bin/GoCart *.go


run:
	./bin/GoCart


all: build run