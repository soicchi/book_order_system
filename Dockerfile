ARG GO_VERSION=1.23.1

FROM golang:$GO_VERSION AS base
RUN curl -L "https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-arm64.tar.gz" | tar xvz && \
    mv ./migrate /usr/bin/migrate

FROM golang:$GO_VERSION AS local
ARG AIR_VERSION=v1.60.0
WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=app/go.mod,target=/app/go.mod \
    --mount=type=bind,source=app/go.sum,target=/app/go.sum \
    go mod download && \
    go install github.com/air-verse/air@$AIR_VERSION

COPY --from=base /usr/bin/migrate /usr/bin/migrate

CMD ["air"]
