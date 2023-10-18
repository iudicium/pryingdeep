FROM golang:1.20.8


WORKDIR /prying-deep

COPY go.mod go.sum ./

RUN go mod download


COPY . ./

EXPOSE 8000

RUN CGO_ENABLED=0 GOOS=linux go build -o pryingdeep ./cmd/pryingdeep
