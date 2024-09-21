FROM golang:1.22-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bank-queue-system ./cmd/http

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/bank-queue-system .

EXPOSE 8080

CMD ["./bank-queue-system"]