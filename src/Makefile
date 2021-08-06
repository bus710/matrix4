# https://github.com/golang/go/wiki/GoArm

build:
ifeq ($(shell uname -p),x86_64)
	go build -race -tags debug -o bin/matrix4-x86_64 cmd/main.go
	chmod 766 bin/matrix4-x86_64
endif
ifeq ($(shell uname -p),aarch64)
	env GOOS=linux GOARCH=arm64 go build -tags debug -o bin/matrix4-arm64 cmd/main.go
	chmod 766 bin/matrix4-arm64
endif
ifeq ($(shell uname -p),arm32)
	env GOOS=linux GOARCH=arm GOARM=7 go build -tags debug -o bin/matrix4-arm cmd/main.go
	chmod 766 bin/matrix4-arm
endif

run: 
ifeq ($(shell uname -p),x86_64)
	./bin/matrix4-x86_64
endif
ifeq ($(shell uname -p),aarch64)
	./bin/matrix4-arm64
endif
ifeq ($(shell uname -p),arm32)
	./bin/matrix4-arm
endif

clean:
	rm -rf bin