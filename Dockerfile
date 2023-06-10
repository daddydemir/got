FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download && go build -o app main.go

EXPOSE 7777

CMD ["./app"]
