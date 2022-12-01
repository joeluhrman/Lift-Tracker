build: 
	go build -o Lift-Tracker.exe

run: build
	./Lift-Tracker.exe

run-prod: build
	./Lift-Tracker.exe -prod

clean:
	go clean
	del *.db

test: 
	go clean -testcache
	go test ./... -v -cover