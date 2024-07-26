build:
	@go build -o ./bin/realesrgan .
run: build
	@LD_LIBRARY_PATH=./realesrgan ./bin/realesrgan
