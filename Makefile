
run:
	docker-compose up

test:
	pytest

lint:
	mypy src/client src/server
	pylint src/client src/server

clean:
	docker-compose rm -f