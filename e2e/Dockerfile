# Bubbler E2E Test Container
#
# Multi-stage build:
#   Stage 1 (builder) — compiles the bubbler binary from Go source
#   Stage 2 (tester)  — installs all language runtimes and runs tests
#
# Usage:
#   docker build -t bubbler-e2e .
#   docker run --rm bubbler-e2e
# or via docker-compose:
#   docker-compose run --rm e2e

# ─────────────────────────────────────────────────────────────────────────────
# Stage 1: Build bubbler binary
# ─────────────────────────────────────────────────────────────────────────────
FROM golang:1.22 AS builder

WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN go build -o /bubbler .

# ─────────────────────────────────────────────────────────────────────────────
# Stage 2: All-runtime test environment
# ─────────────────────────────────────────────────────────────────────────────
FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=UTC

# Install system packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    # Build tools
    gcc g++ make \
    # Java (JDK 17)
    openjdk-17-jdk-headless \
    # Python
    python3 python3-pip \
    # Node.js (from NodeSource 18.x LTS)
    curl ca-certificates gnupg \
    # .NET prerequisites
    libicu70 \
    # Mono (for running .NET Framework 4.7 binaries)
    mono-runtime \
    && rm -rf /var/lib/apt/lists/*

# Node.js 18 LTS
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/*

# .NET 8 SDK
RUN curl -fsSL https://dot.net/v1/dotnet-install.sh | bash -s -- --channel 8.0 \
    && ln -s /root/.dotnet/dotnet /usr/local/bin/dotnet

# Go (for running go tests)
COPY --from=golang:1.22 /usr/local/go /usr/local/go
ENV PATH="/usr/local/go/bin:${PATH}"

# Copy project and compiled bubbler
WORKDIR /app
COPY . .
COPY --from=builder /bubbler /app/bubbler

# Set up Go module cache
ENV GOPATH=/go
ENV PATH="/go/bin:${PATH}"
RUN go mod download

# Make test scripts executable
RUN chmod +x /app/e2e/run_tests.sh \
    /app/e2e/compiler/test_options.sh \
    /app/e2e/compiler/test_errors.sh

ENV BUBBLER=/app/bubbler

CMD ["bash", "/app/e2e/run_tests.sh"]
