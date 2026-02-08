FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

RUN apk add --no-cache git

# Copy only go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY src ./src

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o myapp ./src

# ---- Runtime image ----
FROM alpine:3.20.2

RUN apk add --no-cache ca-certificates

WORKDIR /
COPY --from=builder /app/myapp /myapp

ENTRYPOINT ["/myapp"]
