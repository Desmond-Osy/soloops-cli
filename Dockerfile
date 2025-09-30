# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
ARG VERSION=dev
ARG GIT_COMMIT=none
ARG BUILD_DATE=unknown

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X main.Version=${VERSION} -X main.GitCommit=${GIT_COMMIT} -X main.BuildDate=${BUILD_DATE}" \
    -o soloops ./cmd/soloops

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates terraform

# Create non-root user
RUN addgroup -g 1000 soloops && \
    adduser -D -u 1000 -G soloops soloops

# Set working directory
WORKDIR /workspace

# Copy binary from builder
COPY --from=builder /build/soloops /usr/local/bin/soloops

# Change ownership
RUN chown -R soloops:soloops /workspace

# Switch to non-root user
USER soloops

# Set entrypoint
ENTRYPOINT ["soloops"]

# Default command
CMD ["--help"]