build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o server ./main.go
	./swgspecgen.sh

image: build
	docker build .
