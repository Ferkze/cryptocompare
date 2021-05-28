FROM golang:1.16

WORKDIR /go/src/cryptocompare
COPY . .

RUN go get -d -v ./...
RUN go build -v .

CMD ["./cryptocompare"]