FROM golang:1.23-alpine


WORKDIR /app

RUN apk update && apk add git 

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service

CMD ["./user-service"]
