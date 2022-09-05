.PHONY: run build dockerImage dockerContainer

run:
	go run ./cmd/main.go

build:
	CGO_ENABLED=0 go build -o ./app/api ./cmd/main.go

dockerImage:
	docker image build -f Dockerfile -t correctorImage .

dockerContainer:
	docker container run -p 8081:8081 --detach --name correctorContainer correctorImage

clean:
	docker stop correctorContainer
	docker image rm -f correctorImage
	docker container rm -f correctorContainer