include scripts/go.mk

ifeq ($(OS),Windows_NT)
  include scripts/windows.mk
else
  include scripts/linux.mk
endif

# Define phony targets
.PHONY: build clean run help

# Build the project
build:
	$(GO) build -o $(BINARY) $(MAIN)

# Run the project
run: build
	./$(BINARY)

# Remove generated files
clean:
	-$(RM) $(BINARY)

# Show available commands
help:
	@echo "Available commands:"
	@echo "  make build  - Build the project"
	@echo "  make run    - Build and run the project"
	@echo "  make clean  - Remove the executable"
	@echo "  make help   - Show this help"
