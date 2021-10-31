FROM golang:1.17 as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY apis ./apis

COPY config ./config

COPY internal ./internal

COPY xds/main.go ./main.go

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

FROM alpine:3

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /server

CMD ["/server"]