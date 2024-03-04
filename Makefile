build:
	go build -v ./cmd/invoice-service

protoc:
	protoc proto/invoice.proto --go_out=. --go-grpc_out=.
deploy:
	 docker-compose build && docker-compose down && docker-compose up -d