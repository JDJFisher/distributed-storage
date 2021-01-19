N=1

run:
	docker-compose up --scale server=$(N)

test:
	pytest

lint:
	# mypy client server
	pylint client server