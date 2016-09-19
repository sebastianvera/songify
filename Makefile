APP=spotify
BIN=./bin/$(APP)

all: run

clean c:
	@rm -f $(BIN)

build b: clean
	@go build -o $(BIN)

run r: build
	@$(BIN)

.PHONY: run r clean c all
