N=1

run:
	docker-compose up --scale server=$(N)

test:
	pytest

lint:
	mypy src/client src/server
	pylint src/client src/server

clean:
	docker-compose rm -f