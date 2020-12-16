doc:
	swag init

up:
	make doc
	go build
	./keypass

clean:
	rm keypass
	rm local.db
	rm index.html

report:
	go test -covermode=set -coverprofile=keypass-test.txt ./...
	go tool cover -html=keypass-test.txt -o index.html
