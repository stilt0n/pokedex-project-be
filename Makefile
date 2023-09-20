dev:
	go run ./cmd/api --dev

setup:
	python ./scripts/db_setup.py && echo 'starting docker container...' && cd local-dev-env && docker compose up -d && cd ..

down:
	echo 'shutting down docker container...' && cd local-dev-env && docker compose down && cd ..

nuke:
	git clean -xdf