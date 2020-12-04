.PHONY: build-image
build-image: build-master-server

.PHONY: build-master-server
build-master-server:
	docker build -f build/master/Dockerfile -t k8s-master:v1 .

.PHONY: run-master
run-master:
	go run cmd/master/main.go

.PHONY: build-master
run-master:
	go build -o  cmd/master/master cmd/master/main.go