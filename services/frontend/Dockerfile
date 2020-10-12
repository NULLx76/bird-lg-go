# Build static assets such as CSS
FROM node as assets

WORKDIR /build

COPY services/frontend/package.json .
COPY services/frontend/yarn.lock .
RUN yarn install --frozen-lockfile
COPY services/frontend /build
RUN yarn build
RUN yarn minify

# Build qtc templates
FROM golang as templates
WORKDIR /build

RUN go get -u github.com/valyala/quicktemplate/qtc

# Generate templates
COPY services/frontend/templates templates
RUN qtc -dir=templates

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
COPY --from=templates /build/templates /build/services/frontend/templates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o frontend /build/services/frontend

# Run stage
FROM scratch
WORKDIR /app

COPY --from=build /build/frontend /app/frontend
COPY --from=assets /build/build /app/build

ENTRYPOINT ["/app/frontend"]
