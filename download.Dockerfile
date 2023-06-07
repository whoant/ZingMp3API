FROM golang:1.19-alpine as builder
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -o main cmd/download/*.go

FROM alpine:3.16
WORKDIR /app

COPY --from=builder app/main .

CMD ["/app/main"]
