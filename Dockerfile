FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod (no go.sum since there are no external dependencies yet)
COPY go.mod ./

# Copy the source code to the workspace
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o test_v2_app ./cmd/

# Start a new stage from a minimal alpine image
FROM alpine:latest  

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/test_v2_app .

# Command to run the executable
CMD ["./test_v2_app"]
