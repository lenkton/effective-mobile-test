FROM golang:1.24.5-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/server

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build /api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
