FROM golang:1.22-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod ?go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

FROM alpine:3 AS runner

WORKDIR /req-res

COPY --from=builder /usr/local/bin/app  ./app

CMD ["/req-res/app"]