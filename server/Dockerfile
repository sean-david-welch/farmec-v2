# syntax=docker/dockerfile:1

# Create a stage for building the application.
ARG GO_VERSION=1.22.0
FROM golang:${GO_VERSION} AS build
WORKDIR /src

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application.
ARG TARGETARCH
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

FROM alpine:latest AS final

# Install runtime dependencies.
RUN apk --no-cache add ca-certificates tzdata

# Create a non-privileged user.
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# Copy the executable.
COPY .env /bin/
COPY --from=build /bin/server /bin/

# Expose the port.
EXPOSE 80

# Run the application.
ENTRYPOINT ["/bin/server"]
