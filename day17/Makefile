TARGET:=main.wasm



$(TARGET): main.go
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o ./$@  .
	mv ./$(TARGET) ./build/

.phony:  clean

clean:
	rm -f ./$(TARGET)
	rm -f ./build/$(TARGET)
