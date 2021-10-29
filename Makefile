
distMergeSort-build:
	@echo "Building distMergeSort Docker Image"	
	docker build -t sort -f Dockerfile .

distMergeSort-run:
	@echo "Running Single distMergeSort Docker Container"
	docker run -p 8080:8080 -d sort

provision:
	@echo "Provisioning Sort Cluster"	
	bash scripts/provision.sh 3

info:
	echo "Sort Cluster Nodes"
	docker ps | grep 'sort'
	docker network ls | grep sort_network

clean:
	@echo "Cleaning Sort Cluster"
	bash scripts/teardown.sh

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
