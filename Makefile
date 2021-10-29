
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

golanglintci:
	@echo "golanglintci distMergeSort Server"	
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.42.1 golangci-lint run -v

lint-dockerfile:
	@echo "lint distMergeSort Dockerfile"	
	docker run --rm -i hadolint/hadolint < Dockerfile

test:
	@echo "Testing distMergeSort Server"	
	go test -v --cover ./...
