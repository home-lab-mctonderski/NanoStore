# Stage 1: Build
FROM golang:1.23-alpine3.20 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o storage-service cmd/main.go

# Stage 2: Minimal runtime
FROM scratch
COPY --from=build /app/storage-service /storage-service
ENTRYPOINT ["/storage-service"]
