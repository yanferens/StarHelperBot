APP_NAME=StarHelperBot
MAIN_DIR=cmd/StarHelperBot

.PHONY: all build tidy run fmt clean

build:
	go build -o $(APP_NAME) ./$(MAIN_DIR)

tidy:
	go mod tidy

run: tidy build
	./$(APP_NAME)

fmt:
	gofmt -w .

clean:
	rm -f $(APP_NAME)
