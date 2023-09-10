.PHONY : format install build

#Website/Restfull API

# for run server
run:
	@echo "Running server..."
	go run ./app/main.go

# for development
init:
	@echo "Initializing dependencies..."
	go mod init
	go mod tidy

# for download dependencies
install:
	@echo "Downloading dependencies..."
	go mod download

# for format code
build:
	@echo "building binary..."
	go build ./app/main.go

# for start server
start:
	@echo "Starting server..."
	./app/main

swag:
	@echo "Generating swagger docs..."
	swag init -g ./app/main.go

swag-debug:
	@echo "Generating swagger docs..."
	swag init --parseDependency --parseInternal --parseDepth 2 -g app/main.go -d ./

# for clean binary
clean:
	@echo "Cleaning..."
	rm -rf ./app/main.exe
# live reload using nodemon: npm -g i nodemon
run-nodemon:
	@echo "Running server with nodemon..."
	nodemon --exec go run ./app/main.go