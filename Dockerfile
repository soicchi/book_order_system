ARG GO_VERSION=1.23.1

FROM golang:$GO_VERSION AS local
ARG AIR_VERSION=v1.60.0
WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=app/go.mod,target=/app/go.mod \
    --mount=type=bind,source=app/go.sum,target=/app/go.sum \
    go mod download && \
    go install github.com/air-verse/air@$AIR_VERSION

CMD ["air"]
