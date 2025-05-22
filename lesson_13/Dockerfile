FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/server
RUN go build -o /app/client ./cmd/client

FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache netcat-openbsd

COPY --from=builder /app/server /app/server
COPY --from=builder /app/client /app/client

ARG component=server
ENV COMPONENT=${component}

EXPOSE 8080

RUN echo '#!/bin/sh' > /app/entrypoint.sh && \
    echo 'if [ "$COMPONENT" = "server" ]; then' >> /app/entrypoint.sh && \
    echo '  exec /app/server' >> /app/entrypoint.sh && \
    echo 'else' >> /app/entrypoint.sh && \
    echo '  exec /app/client' >> /app/entrypoint.sh && \
    echo 'fi' >> /app/entrypoint.sh && \
    chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]