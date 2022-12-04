build: 
	go build -o ./Lift-Tracker.exe

runtest: build
	./Lift-Tracker.exe

runprod: build
	./Lift-Tracker.exe -prod

clean:
	go clean

test: 
	go clean -testcache
	go test ./... -v -cover