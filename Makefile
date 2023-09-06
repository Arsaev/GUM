# Variables
GO = go
LDFLAGS = -lm

# Targets
.PHONY: all test build run clean

all: test build run

test:
    $(GO) test ./...

build:
    $(GO) build -o app

run:
    ./app

clean:
    rm -f app