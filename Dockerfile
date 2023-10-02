FROM golang:1.20.8


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /pryingdeep


EXPOSE 8000


CMD ["./pryingdeep"]
