FROM golang:latest

WORKDIR build

COPY . .

RUN go build

CMD ["go", "run", "main.go"]