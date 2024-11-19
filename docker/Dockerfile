FROM golang:1.21 as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod go.sum ./
RUN go mod tidy

COPY ./ ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM debian:bookworm-slim

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    apt-get clean

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

ENTRYPOINT ["/go/src/app/binary"]
