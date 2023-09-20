# About
A backend for a pokedex app. To help me learn Go.

## Development
### Setup
To run in development mode you will need to set up the database. To do this run:
```sh
make setup-local
```
This will:
- prompt you for input to setup a local env file (this will be skipped if you have a env file already)
- run a postgres docker container you can use in local development
- seed the postgres database with some data to simplify development

To close the docker container use:
```sh
make down
```
You can also do this manually by doing:
```sh
cd local-dev-env
docker compose down
```
### Running server
To run the server first follow the setup steps, or if you have already followed them, make sure the docker container is up. Then run:
```sh
make dev
```
This will start the dev server on port `8080`. Whenever you are done you can use `ctrl+c` to kill the server.
### install make
If you don't have `make` you can run:
```sh
sudo apt install make
```
to get it.