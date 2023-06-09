# --- Build Stage ---

FROM golang:1.20.5-alpine3.18 as build

# Copy mod files.
WORKDIR /src

# Copy source files.
COPY go.mod main.go ./
#COPY ["go.mod", "main.go", "./"]

# Build
RUN go build -o hello

# --- Final Stage ---

FROM alpine:3.18

# Add a group and user.
# See https://stackoverflow.com/a/55757473/12429735
RUN addgroup \
    --system \
    --gid "1001" \
    "golang"

RUN adduser \
    --disabled-password \
    --system \
    --home "/app" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "1001" \
    "golang"

# Install tini.
RUN apk add --no-cache tini

# Copy the binary over.
WORKDIR /app
COPY --from=build --chown=golang:golang /src/hello ./

# Run it as a non-privileged user.
USER golang

# Expose port.
EXPOSE 3000

# Run Node.
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/hello"]
