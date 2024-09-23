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

create_migration:
	@make _docker_run CMD='migrate create -ext sql --dir infrastructure/migrations -seq ${NAME}'

migrate_up:
	@make _docker_exec CMD='migrate -database postgres://postgres:postgres@db:5432/book_order_db?sslmode=disable -path infrastructure/migrations up'

migrate_down:
	@make _docker_exec CMD='migrate -database postgres://postgres:postgres@db:5432/book_order_db?sslmode=disable -path infrastructure/migrations down'

migrate_force:
	@make _docker_exec CMD='migrate -database postgres://postgres:postgres@db:5432/book_order_db?sslmode=disable -path infrastructure/migrations force ${VERSION}'
