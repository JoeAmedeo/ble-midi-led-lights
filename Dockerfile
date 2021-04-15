FROM arm64v8/golang:1.16.3 AS builder

WORKDIR /app

COPY . .

RUN go build 

RUN go test

FROM arm64v8/golang:1.16.3 AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]