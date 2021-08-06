# https://github.com/golang/go/wiki/GoArm

build:
ifeq ($(shell uname -m),x86_64)
	go build -race -tags debug -o bin/matrix4-x86_64 cmd/main.go
	chmod 766 bin/matrix4-x86_64
else
ifeq ($(shell uname -m),aarch64)
	env GOOS=linux GOARCH=arm64 go build -tags debug -o bin/matrix4-arm64 cmd/main.go
	chmod 766 bin/matrix4-arm64
else
ifeq ($(shell uname -m),arm32)
	env GOOS=linux GOARCH=arm GOARM=7 go build -tags debug -o bin/matrix4-arm cmd/main.go
	chmod 766 bin/matrix4-arm
endif
endif
endif

run: 
ifeq ($(shell uname -m),x86_64)
	./bin/matrix4-x86_64
else
ifeq ($(shell uname -m),aarch64)
	./bin/matrix4-arm64
else
ifeq ($(shell uname -m),arm32)
	./bin/matrix4-arm
endif
endif
endif

clean:
	rm -rf bin