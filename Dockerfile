FROM golang:latest

WORKDIR /app
COPY . .

RUN go build eshop_api

EXPOSE 9000
ENTRYPOINT ["./eshop_api"]