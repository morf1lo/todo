FROM golang:1.22.2

COPY . .

RUN go mod download

RUN go build -o app ./cmd/main.go

CMD ["./app"]