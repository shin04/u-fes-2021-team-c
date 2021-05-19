FROM golang:1.16.4-alpine3.13 as builder

WORKDIR /go/src

RUN apk update && apk add --no-cache git curl
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY ./backend .

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -o /go/bin/main -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

EXPOSE 8080

ENTRYPOINT ["/app/main"]