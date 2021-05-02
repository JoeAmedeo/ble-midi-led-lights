FROM --platform=linux/arm golang:1.15 AS builder

WORKDIR /app

COPY . .

ENV GOOS=linux\
    GOARCH=arm

RUN go build 

FROM --platform=linux/amd64 golang:1.15 AS tester

WORKDIR /app

COPY . .

RUN go test -v

FROM --platform=linux/arm scratch AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]