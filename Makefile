
run:
	docker-compose up

grpc:
	python3 -m grpc_tools.protoc --python_out=server/server/ --grpc_python_out=server/server/ -I server/protos/ server/protos/*

test:
	pytest

lint:
	mypy client server
	pylint client server

clean:
	docker-compose rm -f