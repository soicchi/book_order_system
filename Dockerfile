ARG GO_VERSION=1.23.1

FROM golang:$GO_VERSION AS base
ARG COBRA_CLI_VERSION=v1.3.0
WORKDIR /opt

RUN --mount=type=bind,source=app/go.mod,target=/opt/go.mod \
    --mount=type=bind,source=app/go.sum,target=/opt/go.sum \
    go mod download && \
    go install github.com/spf13/cobra-cli@$COBRA_CLI_VERSION

FROM golang:$GO_VERSION AS local
ARG AIR_VERSION=v1.60.0
WORKDIR /opt

COPY --from=base /go/pkg/mod /go/pkg/mod
RUN go install github.com/air-verse/air@$AIR_VERSION

CMD ["air", "-c", ".air.toml"]
