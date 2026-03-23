FROM golang:1.25-alpine AS builder

WORKDIR /app

# gen/proto
COPY gen/proto /app/gen/proto

# user-service
COPY services/user-service/go.mod services/user-service/go.sum ./services/user-service/
RUN cd services/user-service && go mod download

COPY services/user-service ./services/user-service

WORKDIR /app/services/user-service
RUN go build -o /bin/user-service ./cmd/

# ─── Final image ─────────────────────────────────────────
FROM alpine:3.19

COPY --from=builder /bin/user-service /bin/user-service

EXPOSE 50051

CMD ["/bin/user-service"]