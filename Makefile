build:
	go build -o Lift-Tracker.exe

run: build
	./Lift-Tracker.exe

clean:
	go clean

test: 
	go test ./... -v -cover