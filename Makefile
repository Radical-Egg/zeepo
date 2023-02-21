TARGET:=target/zeepo

all: build

build:
	GOARCH=arm64 GOOS=darwin go build -o $(TARGET) cmd/zeepo/zeepo.go

.PHONY: build

clean:
	rm -rf target/
run:
	go run cmd/zeepo/zeepo.go