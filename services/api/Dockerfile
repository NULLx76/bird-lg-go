# Build application
FROM golang as build
WORKDIR /build

# Force modules
ENV GO111MODULE=on

# Cache dependencies
COPY go.* ./
RUN go mod download

# Build project
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api /build/services/api

# Run stage
FROM scratch
WORKDIR /app

COPY --from=build /build/api /app/api

ENTRYPOINT ["/app/api"]
