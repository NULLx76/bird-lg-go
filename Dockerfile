FROM golang as build

# Force modules
ENV GO111MODULE=on

WORKDIR /build

# Cache dependencies
COPY go.sum .
COPY go.mod .
RUN go mod download

# Build project
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o lg-frontend ./frontend

# Run stage
FROM scratch
COPY --from=build /build/lg-frontend /frontend
ENTRYPOINT ["/frontend"]
