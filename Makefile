.PHONY: build run docker-up docker-down clean clean-bin

build: 
  go build -o bin/main ./cmd/main

run-main:
  go run main.go

run: build
      ./bin/main

docker-up:
     docker-compose -f docker/docker-compose.yml up -d

docker-down:
     docker-compose -f docker/docker-compose.yml down

clean: 
    cd docker && docker-compose down -v

clean-bin:
    rm -rf bin/