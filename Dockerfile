FROM --platform=linux/arm golang:1.15 AS builder

# RUN apt-get update && apt-get install -y gcc-aarch64-linux-gnu

WORKDIR /app

COPY . .

RUN echo $TARGETPLATFORM
RUN echo $BUILDPLATFORM

ENV GOOS=linux\
    GOARCH=arm

RUN go build 

RUN go test -v

FROM --platform=linux/arm alpine:3.13 AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]