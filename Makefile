
build:
	@echo "Building distMergeSort Server"	
	go build -o bin/distMergeSort main.go

fmt:
	@echo "go fmt distMergeSort Server"	
	go fmt ./...

vet:
	@echo "go vet distMergeSort Server"	
	go vet ./...

lint:
	@echo "go lint distMergeSort Server"	
	golint ./...

test:
	@echo "Testing distMergeSort Server"	
	go test -v --cover ./...
