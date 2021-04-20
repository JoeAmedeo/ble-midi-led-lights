FROM arm32v7/golang:1.15-alpine AS builder

# RUN apt-get update && apt-get install -y gcc-aarch64-linux-gnu

WORKDIR /app

COPY . .

RUN go build 

FROM arm32v7/alpine:3.13 AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]