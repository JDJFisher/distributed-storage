
run:
	docker-compose up

grpc:
	python3 -m grpc_tools.protoc --python_out=src/server/ --grpc_python_out=src/server/ -I src/server/protos/ src/server/protos/*

test:
	pytest

lint:
	mypy src/client src/server
	pylint src/client src/server

clean:
	docker-compose rm -f