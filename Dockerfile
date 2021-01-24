FROM golang:1-stretch AS base

FROM base AS dev

WORKDIR /dev

COPY go.mod ./

RUN go mod download

ENV MODE=development

CMD ["go", "run", "sample.go"]

FROM base AS builder

WORKDIR /build

COPY go.mod ./

RUN go mod download

COPY . ./

ENV MODE=production

RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

FROM alpine:latest AS production

COPY --from=builder /build .

ENTRYPOINT ["./main"]