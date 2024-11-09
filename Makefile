_docker_run:
	docker compose run --rm api ${CMD}

_docker_exec:
	docker compose exec api ${CMD}

docker_build:
	docker compose build

docker_up:
	docker compose up

docker_down:
	docker compose down

docker_stop:
	docker compose stop

go_fmt:
	@make _docker_run CMD='go fmt ./...'

go_vet:
	@make _docker_run CMD='go vet ./...'

go_test:
	@make _docker_run CMD='go test -v -cover ./...'

go_get:
	@make _docker_run CMD='go get ${PKG}'

go_tidy:
	@make _docker_run CMD='go mod tidy'

cobra_add:
	@make _docker_run CMD='cobra-cli add ${NAME}'
