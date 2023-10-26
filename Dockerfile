FROM registry.hub.docker.com/library/golang:1.20.10-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Final stage
FROM scratch as output
COPY --from=builder /app/myapp /myapp
ENTRYPOINT ["/myapp"]
