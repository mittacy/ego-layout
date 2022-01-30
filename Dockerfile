FROM golang:1.16 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn,direct go build -o ./bin/server main.go

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/.env.development /app

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./server"]

CMD ["start", "http", "-c", ".env.development", "-p", "8080", "-e", "production"]