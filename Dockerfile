FROM golang:1.22-alpine3.19
LABEL authors="novo"

COPY webook /app/webook
WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/webook"]