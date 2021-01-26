run:
	docker-compose up

grpc:
	python3 -m grpc_tools.protoc --python_out=server/server/ --grpc_python_out=server/server/ -I server/protos/ server/protos/*

test:
	pytest

clean:
	docker-compose rm -f