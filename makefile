APP_NAME=TempWiFI-Creator

run:
	@echo "Running the application"
	@go run main.go

build:
	@echo "Compiling for every OS and Platform"
	@go build -o bin/$(APP_NAME) main.go