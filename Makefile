TARGET:=target/zeepo

all: build

build:
	GOARCH=arm64 GOOS=darwin go build -trimpath -o $(TARGET) cmd/zeepo/zeepo.go

.PHONY: build

clean:
	rm -rf target/
run:
	go run -trimpath cmd/zeepo/zeepo.go