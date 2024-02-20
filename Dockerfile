FROM cgr.dev/chainguard/go@sha256:e904fb571ca3f60d69e743224f8fc8cfcd0c8faaa29a9307c71e96dde0b70bf1 AS builder

WORKDIR /app
COPY . /app

RUN go install github.com/swaggo/swag/cmd/swag@latest; \
    /root/go/bin/swag init; \
    go mod tidy; \
    go build -o main .

FROM cgr.dev/chainguard/glibc-dynamic@sha256:1790b6f92f7b5d817fd936e95d8a0dfbbc0f5a3648eac67961e66ab5ffbb7fbd

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs docs

ENV ARANGO_HOST localhost
ENV ARANGO_USER root
ENV ARANGO_PASS rootpassword
ENV ARANGO_PORT 8529
ENV MS_PORT 8080

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
