FROM golang:1.21 as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod go.sum ./
RUN go mod tidy

COPY ./ ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM scratch

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

ENTRYPOINT ["/go/src/app/binary"]
