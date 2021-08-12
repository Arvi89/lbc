FROM golang:1.14

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./server ./cmd

CMD [ "/app/server" ]
