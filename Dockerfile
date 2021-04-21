FROM --platform=linux/arm golang:1.15 AS builder

# RUN apt-get update && apt-get install -y gcc-aarch64-linux-gnu

WORKDIR /app

COPY . .

ENV GOOS=linux\
    GOARCH=arm

RUN go build 

RUN go test -v

FROM --platform=linux/arm alpine:3.13 AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]