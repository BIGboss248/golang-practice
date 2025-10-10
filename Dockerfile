FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

# Building the image note go compiler automaticly uses these env variables when compiling
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -ldflags='-s -w' -o /app/main .

FROM alpine:3.18

WORKDIR /app

# Copy the statically-built binary from the builder stage
COPY --from=builder /app/main /app/main

# Ensure the binary is executable
RUN chmod +x /app/main

CMD ["/app/main"]
