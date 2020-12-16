up:
	swag init
	go build
	./keypass

clean:
	rm keypass
	rm local.db