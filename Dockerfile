FROM golang:1.20.4-alpine3.18

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE 8000
ENTRYPOINT  ["./main"]