FROM golang:1.14-alpine as builder

COPY . /app

WORKDIR /app

COPY go.mod .

RUN go mod download

ARG version

RUN go build -ldflags "-X main.Version=${version}" -o /devto-exporter

FROM alpine

COPY --from=builder /devto-exporter /devto-exporter

EXPOSE 2112

CMD /devto-exporter