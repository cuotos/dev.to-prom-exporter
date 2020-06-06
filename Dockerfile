FROM golang:1.14-alpine as builder

COPY . /app

WORKDIR /app

RUN go build -o /devto-exporter

FROM alpine

COPY --from=builder /devto-exporter /devto-exporter

EXPOSE 2112

CMD /devto-exporter