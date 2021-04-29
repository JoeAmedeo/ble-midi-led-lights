FROM --platform=linux/arm golang:1.15 AS builder

# RUN apt-get update && apt-get install -y gcc-aarch64-linux-gnu

WORKDIR /app

COPY . .

RUN apt-get install libc6-armel-cross libc6-dev-armel-cross binutils-arm-linux-gnueabi libncurses5-dev build-essential bison flex libssl-dev bc

ENV GOOS=linux\
    GOARCH=arm\
    CGO_ENABLED=1\
    CC=gcc-arm-linux-gnueabihf

RUN go build 

RUN go test -v

FROM scratch AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]