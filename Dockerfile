FROM golang:1-stretch AS base

FROM base AS dev

ENV MODE=development

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

CMD ["go", "run", "cmd/passport/main.go"]

FROM base AS builder

WORKDIR /build

ENV MODE=production

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o passport cmd/passport/main.go

FROM scratch AS production

WORKDIR /app

COPY --from=builder /build/passport .

ENTRYPOINT ["./passport"]
