TARGET:=target/zeepo

all: $(TARGET)

bookworm:
	go build -o $(TARGET) cmd/zeepo/zeepo.go

clean:
	rm -rf target/
run:
	go run cmd/zeepo/zeepo.go