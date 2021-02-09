FROM alpine
EXPOSE 2112
COPY devto-exporter /devto-exporter
CMD /devto-exporter