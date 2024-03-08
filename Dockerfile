# Etapa de Build
FROM golang:1.20.1-alpine3.17 as builder

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add git make && \
    apk add --no-cache --virtual openssh

COPY . .

RUN go mod vendor && go build -o engine main.go

# Etapa binario
FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app && mkdir /app/configs
WORKDIR /app

EXPOSE 8090
COPY --from=builder /app/engine /app
CMD ["sh", "-c", "/app/engine"]