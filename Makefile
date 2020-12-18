

doc:
	swag init

up:
	make doc
	docker-compose up -d
	go build
	./keypass

clean:
	rm keypass
	rm local.db
	rm index.html

report:
	go test -covermode=set -coverprofile=keypass-test.txt ./...
	go tool cover -html=keypass-test.txt -o index.html

build:
	docker build -t keypass-api:v0.0.1 .

compose:
	make doc
	make build
	docker-compose up