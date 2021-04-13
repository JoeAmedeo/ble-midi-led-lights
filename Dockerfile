FROM golang:1.16.3 AS Builder

WORKDIR /app

COPY . .

RUN go build 

RUN go test

FROM golang:1.16.3 AS Runner

WORKDIR /app

COPY --from=Builder main main

CMD ["./main"]