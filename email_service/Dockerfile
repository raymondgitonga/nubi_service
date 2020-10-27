FROM golang:1.15.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT 8082

RUN go build

CMD ["./email_service"]