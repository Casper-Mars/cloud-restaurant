FROM debian:stable-slim


COPY  ./bin /app
COPY  ./configs /app/configs

WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ["./server", "-conf", "./configs"]
