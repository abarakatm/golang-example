IMAGE_NAME=abarakatm/rafee-hello-world
VERSION=0.0.1


.PHONY: build
build:
	rm -rf build
	go build -o build/hello .

.PHONY: dockerize
dockerize:
	docker build -t ${IMAGE_NAME}:${VERSION} \
    		--platform linux/amd64 -f Dockerfile .

.PHONY: start
start: build
	./build/hello

.PHONY: docker-run
docker-run:
	docker compose up --build

.PHONY: docker-push
docker-push: dockerize
	docker push ${IMAGE_NAME}:${VERSION}



