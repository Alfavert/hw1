FROM golang:latest

WORKDIR build

COPY . .

RUN go build -o main .

CMD ["./main"]