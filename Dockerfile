FROM cgr.dev/chainguard/go@sha256:06a833de4ad9e9b2095de8af3a7097392649ae37528553be73b0855aed425c53 AS builder

WORKDIR /app
COPY . /app

RUN go mod tidy; \
    go build -o main .

FROM cgr.dev/chainguard/glibc-dynamic@sha256:b7eac07361485e71cd0908a6a913f8b2006cf04fef6137f2c9291be00a67ebcc

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
