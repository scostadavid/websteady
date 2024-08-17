FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o tiger ./cmd/app

EXPOSE 8080

CMD [ "./tiger" ]