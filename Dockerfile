FROM golang:1.14

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * ./
RUN go build -o /server

CMD [ "/server" ]
