FROM golang:1.16 AS builder

# RUN apt-get update && apt-get install -y gcc-aarch64-linux-gnu

WORKDIR /app

COPY . .

RUN go build 

RUN go test -v

FROM alpine:3.13 AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]