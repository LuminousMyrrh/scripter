main:
	go build -o dist/scripter cmd/scripter/main.go

clean:
	rm -f $(TARGET)
	mkdir dist
