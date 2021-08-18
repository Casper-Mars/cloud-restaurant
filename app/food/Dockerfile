FROM golang:1.16.7-buster AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim


COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /app/configs

WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ["./server", "-conf", "./configs"]
