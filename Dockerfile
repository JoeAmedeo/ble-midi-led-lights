FROM --platform=linux/arm golang:1.15 AS builder

WORKDIR /app

COPY . .

ENV GOOS=linux\
    GOARCH=arm

RUN go build 

RUN go test *.go -v

FROM --platform=linux/arm scratch AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]