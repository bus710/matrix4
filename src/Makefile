build:
	go build -race -tags debug -o bin/matrix4-x86_64 cmd/main.go
	chmod 766 bin/matrix4-x86_64

# https://github.com/golang/go/wiki/GoArm
build-arm32:
	env GOOS=linux GOARCH=arm GOARM=7 go build -tags debug -o bin/matrix4-arm cmd/main.go
	chmod 766 bin/matrix4-arm

build-arm64:
	env GOOS=linux GOARCH=arm64 go build -tags debug -o bin/matrix4-arm64 cmd/main.go
	chmod 766 bin/matrix4-arm64

run: 
	./bin/matrix4-x86_64

run-arm32: 
	./bin/matrix4-arm

run-arm64:
	./bin/matrix4-arm64

clean:
	rm -rf bin
