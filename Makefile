IMB_BINARY = imb
MONKEY_BINARY = monkey

IMB_SRC = ./cmds/imb/main.go
MONKEY_SRC = ./cmds/monkey/main.go

all: $(IMB_BINARY) $(MONKEY_BINARY)

$(IMB_BINARY): $(IMB_SRC)
	@echo "Building REPL binary..."
	go build -o $(IMB_BINARY) $(IMB_SRC)

$(MONKEY_BINARY): $(MONKEY_SRC)
	@echo "Building file execution binary..."
	go build -o $(MONKEY_BINARY) $(MONKEY_SRC)

clean:
	@echo "Cleaning up..."
	rm -f $(IMB_BINARY) $(MONKEY_BINARY)

.PHONY: all clean

