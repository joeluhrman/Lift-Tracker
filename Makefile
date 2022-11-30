build: 
	go build -o Lift-Tracker.exe

run: build
	./Lift-Tracker.exe

clean:
	go clean

cleandb:
	del *.db

cleanall: clean cleandb

test: 
	go test ./... -v -cover