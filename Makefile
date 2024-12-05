_docker_exec:
	docker compose exec api ${CMD}

_docker_run_test:
	docker compose run --rm test_api ${CMD}

docker_build:
	docker compose build

docker_up:
	docker compose up

docker_down:
	docker compose down

docker_stop:
	docker compose stop

go_fmt:
	@make _docker_exec CMD='go fmt ./...'

go_vet:
	@make _docker_exec CMD='go vet ./...'

go_test:
	@make _docker_run_test CMD='go test -v -cover ./...'

go_get:
	@make _docker_exec CMD='go get ${PKG}'

go_tidy:
	@make _docker_exec CMD='go mod tidy'

cobra_add:
	@make _docker_exec CMD='cobra-cli add ${NAME}'

migrate_up:
	@make _docker_exec CMD='go run main.go migrateup'

make_migration:
	@make _docker_exec CMD='go run main.go makemigration -n ${NAME}'
