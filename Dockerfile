ARG GO_VERSION=1.23.1

FROM golang:$GO_VERSION AS base
ARG COBRA_CLI_VERSION=v1.3.0
WORKDIR /opt

RUN --mount=type=bind,source=app/go.mod,target=/opt/go.mod \
    --mount=type=bind,source=app/go.sum,target=/opt/go.sum \
    go mod download

FROM golang:$GO_VERSION AS local
ARG AIR_VERSION=v1.61.1
WORKDIR /opt

RUN go install github.com/air-verse/air@$AIR_VERSION

COPY --from=base /go/pkg/mod /go/pkg/mod

CMD ["air", "c", ".air.toml"]
