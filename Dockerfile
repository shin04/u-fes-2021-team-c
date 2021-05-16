FROM golang:1.16.4-alpine3.13 as builder

RUN apk update \
    && apk add --no-cache git curl \
    && chmod +x ${GOPATH}/bin/air

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY ./backend .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./cmd

FROM alpine:3.12

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main web"]