build: 
	go build -o ./Lift-Tracker.exe

run: build
	./Lift-Tracker.exe

clean:
	go clean

test: 
	go clean -testcache
	go test ./... -cover

seed:
	go build -o ./seeder.exe ./seeder/seeder.go
	./seeder.exe