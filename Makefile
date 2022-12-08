build: 
	go build -o ./Lift-Tracker.exe

runtest: build
	./Lift-Tracker.exe 

clean:
	go clean

test: 
	go clean -testcache
	go test ./... -v -cover