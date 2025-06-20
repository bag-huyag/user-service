FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/main.go

FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=builder /app/user-service .

USER nonroot:nonroot

EXPOSE 50052

ENTRYPOINT ["/app/user-service"]
