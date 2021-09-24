## Build api
FROM golang:1.16

WORKDIR /periskop-pushgateway

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o app .

ENV PORT 8080

CMD ["/periskop-pushgateway/app"]
