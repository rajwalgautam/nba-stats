FROM golang:1.25-alpine AS builder

WORKDIR /app

# Cache and download dependencies from mod file
COPY go.mod go.sum ./
RUN go mod download

# Copy dependencies and source code
COPY . .

# Install build dependencies (gcc and g++) - we need gcc for our sql driver
RUN apk --no-cache add gcc g++ sqlite

# Build the Go binary
RUN CGO_ENABLED=1 GOOS=linux go build -trimpath -ldflags="-s -w" -o gamestatsjob ./cmd/gamestatsjob 

# Second stage - create small image that the Go binary will run in
FROM alpine:latest

WORKDIR /app

# Install timezone data and certificates
# We will eventually need this for TLS and timezone support 
RUN apk add --no-cache ca-certificates tzdata

# Create non-root user for security
RUN addgroup -S appuser \
 && adduser -S -G appuser -H -s /sbin/nologin appuser

# Copy the Go binary from the builder stage
COPY --from=builder /app/gamestatsjob .

# Use non-root user to run our application for security
USER appuser

# Run the Go binary
ENTRYPOINT ["/app/gamestatsjob"]