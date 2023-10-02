FROM golang:1.20.8


WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download


COPY . ./


RUN CGO_ENABLED=0 GOOS=linux go build -o pryingdeep ./cmd/prying-deep


ENTRYPOINT [ "./pryingdeep" ]
