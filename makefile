
build:
	go build -o bin/GoCart *.go

Docker-build:
	GOOS=linux CGO_ENABLED=0 go build -o main *.go

run:
	./bin/GoCart



docker-compose-up:
	sudo docker compose up --build


docker-compose-down:
	sudo docker compose down

# make docker to start docker

docker: Docker-build docker-compose-up


all: build run