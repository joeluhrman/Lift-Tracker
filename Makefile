build: clean
	go build -o Lift-Tracker.exe

run: clean build
	./Lift-Tracker.exe

clean:
	go clean

test: 
	go test ./... -v -cover