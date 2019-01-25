DOCKER_IMAGE_NAME=rsmitty/pkg-lister
VERSION ?= $(shell bash -c 'read -p "Enter version tag: " pwd; echo $$pwd')

all: linux docker

darwin:
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ../bin/pkg-lister-darwin ../*.go

linux:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../bin/pkg-lister-linux ../*.go

docker: docker-build docker-push

docker-build: 
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) -f ./Dockerfile ../

docker-push:
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)

clean:
	rm -rf ../bin
