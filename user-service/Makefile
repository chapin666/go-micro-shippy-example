build:
		protoc -I . --go_out=plugins=micro:. proto/user/user.proto
		GOOS=linux GOARCH=amd64 go build
		docker build -t user-service .
run:
		docker run -p 50053:50051 \
			-e MICRO_SERVER_ADDRESS=:50051 \
			-e MICRO_REGISTRY=mdns	\
			-e MICRO_BROKER=nats \
			-e MICRO_BROKER_ADDRESS=0.0.0.0:4222 \
			user-servcie