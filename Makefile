
install:
	pip3 install -r requirements.txt

test:
	pytest

lint:
	mypy main
	pylint main
