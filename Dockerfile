FROM golang:1.22-alpine AS builder
WORKDIR /usr/src/app
COPY go.mod ?go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/go-echo-server ./...

FROM alpine:3 AS runner
WORKDIR /app
COPY --from=builder /usr/local/bin/go-echo-server  ./go-echo-server
CMD ["/app/go-echo-server"]