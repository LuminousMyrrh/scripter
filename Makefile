main:
	go build -o dist/main cmd/scripter/main.go

clean:
	rm -f $(TARGET)
	mkdir dist
